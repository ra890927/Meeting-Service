import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';


import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgIf } from '@angular/common';
import { FooterComponent } from '../layout/footer/footer.component';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    MatInputModule,
    MatFormFieldModule,
    MatSelectModule,
    MatButtonModule,
    ReactiveFormsModule,
    FooterComponent,
    NgIf
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  constructor(private router: Router) { }

  @ViewChild("userNameInput")
  userNameInput!: ElementRef<MatInput>;
  @ViewChild("passwordInput")
  passwordInput!: ElementRef<MatInput>;

  ngOnInit(): void {
  }

  // password validation
  password = new FormControl('', [
    Validators.required,
    Validators.minLength(8),
    Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]+$')
  ]);

  getErrorMessage() {
    if (this.password.hasError('required')) {
      return 'You must enter a value';
    }
    else if (this.password.hasError('minlength')) {
      return 'Must be at least 8 characters';
    }

    return this.password.hasError('pattern') ? 'Must contain at least 1 letter and 1 number' : '';
  };
  
  navigate(path: string) {
    this.router.navigate([path]); // navigate to the path
  }

}
