import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

import { MatButtonModule } from '@angular/material/button';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { FormControl, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    MatButtonModule,
    MatInputModule,
    MatFormFieldModule,
    MatIconModule,
    FormsModule,
    ReactiveFormsModule,
    NgIf
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {

  @ViewChild("userNameInput")
  userNameInput!: ElementRef<MatInput>;
  @ViewChild("emailInput")
  emailInput!: ElementRef<MatInput>;
  @ViewChild("passwordInput")
  passwordInput!: ElementRef<MatInput>;
  @ViewChild("confirmPasswordInput")
  confirmPasswordInput!: ElementRef<MatInput>;


  hidePassword = true;
  hideConfirmPassword = true;
  constructor(private router: Router) { }

  ngOnInit(): void {
  }

  // navigate to the login page
  navigate(path: string) {
    this.router.navigate([path]);
  }

  // email validation
  email = new FormControl('', [Validators.required, Validators.email]);

  getErrorMessage() {
    if (this.email.hasError('required')) {
      return 'You must enter a value';
    }

    return this.email.hasError('email') ? 'Not a valid email' : '';
  }

  // password validation
  password = new FormControl('', [
    Validators.required,
    Validators.minLength(8),
    Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]+$')
  ]);

  getErrorMessagePassword() {
    if (this.password.hasError('required')) {
      return 'You must enter a value';
    }
    else if (this.password.hasError('minlength')) {
      return 'Must be at least 8 characters';
    }

    return this.password.hasError('pattern') ? 'Must contain at least 1 letter and 1 number' : '';
  };

  // confirm password validation xxx
  getErrorMessageConfirmPassword() {
    if (this.confirmPasswordInput.nativeElement.value != this.passwordInput.nativeElement.value) {
      return 'Passwords do not match';
    }
    return null;
  };

  // rest the form
  reset(){
    this.userNameInput.nativeElement.value = '';
    this.emailInput.nativeElement.value = '';
    this.passwordInput.nativeElement.value = '';
    this.confirmPasswordInput.nativeElement.value = '';

  }

}
