import { Component, ElementRef, ViewChild } from '@angular/core';
import { users } from './users';

import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';

import { HeaderComponent } from '../layout/header/header.component';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatInput } from '@angular/material/input';

@Component({
  selector: 'app-monitor',
  standalone: true,
  imports: [
    HeaderComponent,
    MatTabsModule,
    MatCardModule,
    MatListModule,
    MatIconModule,
    CommonModule

  ],
  templateUrl: './monitor.component.html',
  styleUrl: './monitor.component.css'
})
export class MonitorComponent {
  usersList: users[] = [];
  usersEditing: users | undefined;
  isEditing: boolean = false;

  @ViewChild("userNameInput")
  userNameInput!: ElementRef<MatInput>;
  @ViewChild("emailInput")
  emailInput!: ElementRef<MatInput>;
  @ViewChild("roleInput")
  roleInput!: ElementRef<MatInput>;
  @ViewChild("detailsInput")
  detailsInput!: ElementRef<MatInput>;

  ngOnInit(): void {
    this.usersList.push({
      id: '001',
      userName: 'John Doe',
      email: 'pat@example.com',
      role: 'Admin',
      details: 'This is a test user'
    });
    const usersJson = localStorage.getItem("userslist");
    if (usersJson) this.usersList = JSON.parse(usersJson);
  }


  changeStatus(users: users): void {
    if (this.isEditing){
      // users.status = !users.status;
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
      localStorage.setItem("userslist", JSON.stringify(this.usersList)); // save to local storage
    }
  }

  delete(users: users): void {
    this.usersList = this.usersList.filter(t => t.id !== users.id);
    
    localStorage.setItem("userslist", JSON.stringify(this.usersList));
  }

  edit(users: users): void {
    this.isEditing = !this.isEditing;
    this.usersEditing = users;
    // this.userNameInput.nativeElement.value = users.userName;
    // console.log(this.userNameInput.nativeElement.value);

    // 确保 input 已经被渲染后再设置值
    setTimeout(() => {
      if (this.userNameInput && this.emailInput && this.detailsInput) {
        this.userNameInput.nativeElement.value = users.userName;
        this.userNameInput.nativeElement.focus();
        this.emailInput.nativeElement.value = users.email;
        this.detailsInput.nativeElement.value = users.details;
      }
    }, 0);
    
    localStorage.setItem("userslist", JSON.stringify(this.usersList));
  }

  save(): void {
    if (this.usersEditing) {
      this.usersEditing.userName = this.userNameInput.nativeElement.value;
      this.usersEditing.email = this.emailInput.nativeElement.value;
      this.usersEditing.details = this.detailsInput.nativeElement.value;
      localStorage.setItem("usersList", JSON.stringify(this.usersList));
    }
    this.isEditing = false;
    this.usersEditing = undefined;
  }

}
