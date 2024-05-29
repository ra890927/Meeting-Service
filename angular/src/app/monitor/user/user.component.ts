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

import { UserService } from '../../API/user.service';
import { ItemService } from '../../API/item.service';
import { AdminService } from '../../API/admin.service';


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
  usersEditing: users = {
    id: 0,
    userName: '',
    email: '',
    role: 'user'
  };
  isEditing: boolean = false;
  userNameControl = new FormControl();
  userFromBackend: any;
  connectionError = false;

  constructor(private userService: UserService, private itemservice:ItemService, private adminService: AdminService) {
  }

  ngOnInit(): void {
    this.usersList.push({
      id: 1,
      userName: 'John Doe',
      email: 'pat@example.com',
      role: 'admin'
    });
    this.usersList.push({
      id: 2,
      userName: 'Amy',
      email: 'amy@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 3,
      userName: 'Bob',
      email: 'Bob@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 1,
      userName: 'John Doe',
      email: 'pat@example.com',
      role: 'admin'
    });
    this.usersList.push({
      id: 2,
      userName: 'Amy',
      email: 'amy@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 3,
      userName: 'Bob',
      email: 'Bob@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 1,
      userName: 'John Doe',
      email: 'pat@example.com',
      role: 'admin'
    });
    this.usersList.push({
      id: 2,
      userName: 'Amy',
      email: 'amy@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 3,
      userName: 'Bob',
      email: 'Bob@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 1,
      userName: 'John Doe',
      email: 'pat@example.com',
      role: 'admin'
    });
    this.usersList.push({
      id: 2,
      userName: 'Amy',
      email: 'amy@example.com',
      role: 'user'
    });
    this.usersList.push({
      id: 3,
      userName: 'Bob',
      email: 'Bob@example.com',
      role: 'user'
    });

    // get all users from backend
    this.itemservice.getAllUsers().subscribe((response)=>{
      console.log(response.data.users);
      this.usersList = response.data.users.map((user: any) => {
        return {
          id: user.id,
          email: user.email,
          role: user.role,
          userName: user.username,
        };
      });
    });
  }


  changeStatus(users: users): void {
    if (this.isEditing){
      switch (users.role) {
        case 'admin':
            users.role = 'user';
            this.usersEditing.role = 'user';
            break;
        case 'user':
            users.role = 'admin';
            this.usersEditing.role = 'admin';
            break;
        default:
          users.role = 'user'; // Default to User
          this.usersEditing.role = 'user';
      }

    }
  }

  delete(users: users): void {
    this.usersList = this.usersList.filter(t => t.id !== users.id);
  }

  edit(users: users): void {
    this.isEditing = !this.isEditing;
    this.usersEditing = users;
    this.userNameControl.setValue(users.userName);
  }

  save(): void {
    if(this.usersEditing){
      console.log(this.usersEditing.role);

      this.adminService.updateUser(this.usersEditing.id, this.userNameControl.value, this.usersEditing.email, this.usersEditing.role, '').subscribe(
        (res) => {
          if (res.status === 'success') {
            console.log(this.usersEditing.role);
            const index = this.usersList.findIndex(user => user.id === this.usersEditing.id);
            if (index !== -1) {
              this.usersList[index].userName = this.userNameControl.value;
              this.usersList[index].role = this.usersEditing.role;
              this.isEditing = false;
              this.usersEditing = {
                id: 0,
                userName: '',
                email: '',
                role: 'user'
              };
              this.userNameControl.setValue('');
              console.log("userlists",this.usersList);
            }
            console.log("!");
            
          }else{
            this.isEditing = false;
            this.usersEditing = {
              id: 0,
              userName: '',
              email: '',
              role: 'user'
            };
            this.userNameControl.setValue('');
            console.log('Update failed');
            return
          }
        },
        (error) => {
            console.error('A connection error occurred:', error);
            this.connectionError = true; 
        }
        );
    }


    
  }

}
