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

  @ViewChild("userNameInput")
  userNameInput!: ElementRef<MatInput>;

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
    // users.status = !users.status;
    switch (users.role) {
      case 'Admin':
          users.role = 'Admin';
          break;
      case 'User':
          users.role = 'User';
          break;
      default:
        users.role = 'User'; // Default to incomplete if something goes wrong
  }

    localStorage.setItem("userslist", JSON.stringify(this.usersList)); // save to local storage
  }

  delete(users: users): void {
    this.usersList = this.usersList.filter(t => t.id !== users.id);
    
    localStorage.setItem("userslist", JSON.stringify(this.usersList));
  }

  edit(users: users): void {
    this.usersEditing = users;
    this.userNameInput.nativeElement.value = users.userName;

  
    
    localStorage.setItem("userslist", JSON.stringify(this.usersList));
  }

}
