import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

import { MatButtonModule } from '@angular/material/button';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { FormControl, Validators, FormsModule, ReactiveFormsModule, FormGroup, FormBuilder } from '@angular/forms';
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

  registerForm: FormGroup;

  constructor(private fb: FormBuilder, private router: Router) {
    this.registerForm = this.fb.group({
      userName: ['', [Validators.required, Validators.minLength(4)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [
        Validators.required,
        Validators.minLength(8),
        Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]+$')
      ]],
      confirmPassword: ['', [Validators.required]]
    }, { validators: this.passwordMatchValidator });
  }


  hidePassword = true;
  hideConfirmPassword = true;

  ngOnInit(): void {
  }

  passwordMatchValidator(form: FormGroup) : {[key: string]: any} | null{
    const password = form.get('password');
    const confirmPassword = form.get('confirmPassword');
    return password && confirmPassword && password.value === confirmPassword.value ? null : { 'mismatch': true };
  }

  getErrorMessage(field: string) {
    let control = this.registerForm.get(field);
    if (control && control.errors?.['required']) {
      return 'You must enter a value';
    } else if (control && control.errors?.['minlength'] && field === 'userName') {
      return 'Must be at least 4 characters';
    } else if (control && control.errors?.['email'] && field === 'email') {
      return 'Not a valid email';
    } else if (control && control.errors?.['minlength'] && field === 'password') {
      return 'Must be at least 8 characters';
    } else if (control && control.errors?.['pattern'] && field === 'password') {
      return 'Must contain at least 1 letter and 1 number';
    } else if (control && control.errors?.['mismatch'] && field === 'confirmPassword') {
      return 'Passwords do not match';
    }
    return '';
  }


  // rest the form
  reset(){
    this.registerForm.reset();
  }

  // navigate to the login page
  navigate(path: string) {
    this.router.navigate([path]);
  }

}
