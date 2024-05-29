import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';


import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { AbstractControl,FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgIf } from '@angular/common';
import { AuthService } from '../API/auth.service';
import { co } from '@fullcalendar/core/internal-common';
import { UserService } from '../API/user.service';
@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    MatInputModule,
    MatFormFieldModule,
    MatSelectModule,
    MatButtonModule,
    ReactiveFormsModule,
    NgIf
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  constructor(private router: Router, private authService: AuthService, private userservice: UserService ) { }

  @ViewChild("userNameInput")
  userNameInput!: ElementRef<MatInput>;
  @ViewChild("passwordInput")
  passwordInput!: ElementRef<MatInput>;
  submitted = false;
  loginError = false;
  connectionError = false;
  ngOnInit(): void {
    if (window.sessionStorage.getItem('token')) {
      this.navigate('/profile');//navigate to home page
    }
  }
  //reactive form
  // username validation
  userName = new FormControl('', [
    Validators.required,
    Validators.minLength(4),
  ]);

  // password validation
  password = new FormControl('', [
    Validators.required,
  ]);

  formData = new FormGroup({
    userName: this.userName,
    password: this.password
  });

  get f(): { [key: string]: AbstractControl } {
    return this.formData.controls;
  }
  // login function
  login() {
    this.loginError = false;
    this.connectionError = false;
    if (this.formData.valid) {
      const {userName, password} = this.formData.value;

      if(userName && password){
        this.authService.login(userName, password).subscribe(
        (res) => {
          if (res.status === 'success') {
            //set the token and user details in local storage
            window.sessionStorage.setItem('token', res.data.token);
            this.userservice.saveUser(res.data.user);
            this.navigate('/profile');//navigate to profile page
          }else{
            console.log('Login failed');
            this.loginError = true;
            return
          }
        },
        (error) => {
            console.error('An error occurred:', error);
            this.connectionError = true; // set connection error to true
        }
        );
      } 
    }else{
      this.submitted = true;
      console.log('Invalid form');
    }
  }

  
  navigate(path: string) {
    this.router.navigate([path]); // navigate to the path
  }

}
