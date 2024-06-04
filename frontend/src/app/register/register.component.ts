import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

import { MatButtonModule } from '@angular/material/button';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { FormControl, Validators, FormsModule, ReactiveFormsModule, FormGroup, FormBuilder, ValidatorFn, AbstractControl, ValidationErrors } from '@angular/forms';
import { NgIf } from '@angular/common';
import { FooterComponent } from '../layout/footer/footer.component';
import { UserService } from '../API/user.service';
import { co } from '@fullcalendar/core/internal-common';

// }


export default class Validation {
  static match(controlName: string, checkControlName: string): ValidatorFn {
    return (controls: AbstractControl) => {
      const control = controls.get(controlName);
      const checkControl = controls.get(checkControlName);

      if (checkControl?.errors && !checkControl.errors['matching']) {
        return null;
      }

      if (control?.value !== checkControl?.value) {
        controls.get(checkControlName)?.setErrors({ matching: true });
        return { matching: true };
      } else {
        return null;
      }
    };
  }
}

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
    NgIf,
    FooterComponent
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

  constructor(private fb: FormBuilder, private router: Router, private userService: UserService) {

    this.registerForm = new FormGroup({
      userName: new FormControl('', [Validators.required, Validators.minLength(4)]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(8),
        Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]+$')
      ]),
      confirmPassword: new FormControl('', [Validators.required]),
    }, {validators: [Validation.match('password', 'confirmPassword')]});
  }
  hidePassword = true;
  hideConfirmPassword = true;

  ngOnInit(): void {
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
    } else if (control && control.errors?.['matching'] && field === 'confirmPassword') {
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

  // submit the form
  submit(){
    if (this.registerForm.valid) {
      console.log('Form Submitted!', this.registerForm.value.userName);
      this.userService.register(this.registerForm.value.userName, this.registerForm.value.email, this.registerForm.value.password).subscribe(
        (res) => {
          console.log('res:', res);
          if (res.status === 'success') {
            this.router.navigate(['/login']);
          } else {
            console.log('Register failed');
            return;
          }
        },
        (error) => {
          console.error('A connection error occurred:', error);
          return;
        }
      );
    }

  }

}
