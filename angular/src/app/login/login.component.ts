import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';


import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgIf } from '@angular/common';
import { AuthService } from '../API/auth.service';
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
  constructor(private router: Router, private authService: AuthService ) { }

  @ViewChild("userNameInput")
  userNameInput!: ElementRef<MatInput>;
  @ViewChild("passwordInput")
  passwordInput!: ElementRef<MatInput>;

  ngOnInit(): void {
  }
  //reactive form
  // username validation
  userName = new FormControl('', [
    Validators.required,
    Validators.minLength(4),
    Validators.pattern('^[a-zA-Z0-9]+$')
  ]);

  // password validation
  password = new FormControl('', [
    Validators.required,
    Validators.minLength(8),
    Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]+$')
  ]);

  formData = new FormGroup({
    userName: this.userName,
    password: this.password
  });

  // login function
  login() {
    if (this.formData.valid) {
      const {userName, password} = this.formData.value;
      if(userName && password){
        this.authService.login(userName, password).subscribe((res) => {
          if (res.status === '200') {
            //set the token and user details in local storage
            localStorage.setItem('token', res.body.token);
            localStorage.setItem('user', JSON.stringify(res.body.user));
            this.navigate('/home');//navigate to home page
          }
        });
      } 
    }
  }

  logout() {
    this.authService.logout().subscribe((res) => {
      if (res.status === '200') {
        //clear the token and user details from local storage
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        this.navigate('/login');//navigate to login page
      }
    });
  }
  
  navigate(path: string) {
    this.router.navigate([path]); // navigate to the path
  }

}
