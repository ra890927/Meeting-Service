import { Component, ElementRef, ViewChild } from '@angular/core';
import { users } from '../users';

import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatInput, MatInputModule } from '@angular/material/input';
import { CommonModule } from '@angular/common';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';


@Component({
  selector: 'app-user',
  standalone: true,
  imports: [
    MatTabsModule,
    MatCardModule,
    MatListModule,
    MatIconModule,
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule
    
  ],
  templateUrl: './user.component.html',
  styleUrl: './user.component.css'
})

export class UserComponent {
  usersList: users[] = [];
  usersEditing: users | undefined;
  isEditing: boolean = false;

  userNameControl = new FormControl();
  emailControl = new FormControl();
  detailsControl = new FormControl();

  ngOnInit(): void {
    this.usersList.push({
      id: '001',
      userName: 'John Doe',
      email: 'pat@example.com',
      role: 'Admin',
      details: 'This is a test user'
    });
    const usersJson = localStorage.getItem("usersList");
    if (usersJson) this.usersList = JSON.parse(usersJson);
  }


  changeStatus(users: users): void {
    if (this.isEditing){
      switch (users.role) {
        case 'Admin':
            users.role = 'User';
            break;
        case 'User':
            users.role = 'Admin';
            break;
        default:
          users.role = 'User'; // Default to User
      }
      localStorage.setItem("usersList", JSON.stringify(this.usersList)); // save to local storage
    }
  }

  delete(users: users): void {
    this.usersList = this.usersList.filter(t => t.id !== users.id);
    
    localStorage.setItem("usersList", JSON.stringify(this.usersList));
  }

  edit(users: users): void {
    this.isEditing = !this.isEditing;
    this.usersEditing = users;
    this.userNameControl.setValue(users.userName);
    this.emailControl.setValue(users.email);
    this.detailsControl.setValue(users.details);
    
    localStorage.setItem("usersList", JSON.stringify(this.usersList));
  }

  save(): void {
    if (this.usersEditing) {
      this.usersEditing.userName = this.userNameControl.value;
      this.usersEditing.email = this.emailControl.value;
      this.usersEditing.details = this.detailsControl.value;
      localStorage.setItem("usersList", JSON.stringify(this.usersList));
    }
    this.isEditing = false;
    this.usersEditing = undefined;
    this.userNameControl.setValue('');
    this.emailControl.setValue('');
    this.detailsControl.setValue('');
  }

}
