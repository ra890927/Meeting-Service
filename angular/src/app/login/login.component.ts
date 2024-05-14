import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';


import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    MatInputModule,
    MatFormFieldModule,
    MatSelectModule,
    MatButtonModule
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
  navigate(path: string) {
    this.router.navigate([path]); // navigate to the path
  }

}
