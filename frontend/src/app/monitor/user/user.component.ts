import { Component, ElementRef, ViewChild } from '@angular/core';
import { users } from '../users';

import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatInput, MatInputModule } from '@angular/material/input';
import { AsyncPipe, CommonModule } from '@angular/common';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';

import { UserService } from '../../API/user.service';
import { ItemService } from '../../API/item.service';
import { AdminService } from '../../API/admin.service';
import { MatTooltipModule } from '@angular/material/tooltip';
import { Observable, map, startWith } from 'rxjs';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { co } from '@fullcalendar/core/internal-common';


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
    MatInputModule,
    MatTooltipModule,
    MatAutocompleteModule,
    AsyncPipe
    
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

  filteredOptions: Observable<users[]> | undefined;
  emailSearchControl = new FormControl('');

  constructor(private userService: UserService, private itemservice:ItemService, private adminService: AdminService) {
  }

  ngOnInit(): void {
    // fake data
    // this.usersList = [
    //   {
    //     id: 1,
    //     userName: 'John Doe',
    //     email: '',
    //     role: 'admin'
    //   },
    //   {
    //     id: 2,
    //     userName: 'Jane Doe',
    //     email: '',
    //     role: 'user'
    //   },
    //   {
    //     id: 3,
    //     userName: 'John Smith',
    //     email: '',
    //     role: 'user'
    //   },
    //   {
    //     id: 4,
    //     userName: 'Jane Smith',
    //     email: '',
    //     role: 'user'
    //   },
    //   {
    //     id: 5,
    //     userName: 'John Brown',
    //     email: '',
    //     role: 'user'
    //   },
    //   {
    //     id: 6,
    //     userName: 'Jane Brown',
    //     email: '',
    //     role: 'user'
    //   }];


    // get all users from backend
    this.itemservice.getAllUsers().subscribe((response)=>{
      this.usersList = response.data.users.map((user: any) => {
        return {
          id: user.id,
          email: user.email,
          role: user.role,
          userName: user.username,
        };
      });

      this.filteredOptions = this.emailSearchControl.valueChanges.pipe(
        startWith(''),
        map(value => this._filter(value || '')),
      );
    });

    
  }

  private _filter(value: string): users[] {
    const filterValue = value.toLowerCase();
    const userEmailArray: users[] = this.usersList;
    return this.usersList.filter(userEmailArray => userEmailArray.email.toLowerCase().includes(filterValue));
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
    console.log("filteredOptions", this.filteredOptions);
  }

  save(): void {
    if(this.usersEditing){

      this.adminService.updateUser(this.usersEditing.id, this.userNameControl.value, this.usersEditing.email, this.usersEditing.role, '').subscribe(
        (res) => {
          if (res.status === 'success') {
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
            }
            
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
