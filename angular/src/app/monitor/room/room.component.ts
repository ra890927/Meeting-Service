import { CommonModule } from '@angular/common';
import { rooms, allTags } from '../users';
import { Component, Inject, OnInit } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { MatChipsModule } from '@angular/material/chips';
import { v4 as uuidv4 } from 'uuid'; // generate random id
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms'
import { MatMenuModule } from '@angular/material/menu';

import { ItemService } from '../../API/item.service';
import { AdminService } from '../../API/admin.service';
import { cA } from '@fullcalendar/core/internal-common';

@Component({
  selector: 'app-room',
  standalone: true,
  imports: [
    MatCardModule,
    MatListModule,
    MatIconModule,
    MatFormFieldModule,
    MatDialogModule,
    MatInputModule,
    MatButtonModule,
    MatChipsModule,
    MatMenuModule,
    ReactiveFormsModule,
    CommonModule,
    FormsModule
  ],
  templateUrl: './room.component.html',
  styleUrl: './room.component.css'
})
export class RoomComponent implements OnInit{
  
  roomsList: rooms[] = [{ 
    id: 1,
    roomNumber: 'lab639',
    tags: ['Food Allowed'],
    capacity: 10,
  },
  { id: 2,
    roomNumber: 'lab637',
    tags: ['Food Allowed', 'Projector Available', 'Free WiFi'],
    capacity: 30,
  }];
  roomsEditing: rooms | undefined;
  isEditing: boolean = false;
  roomNumberControl = new FormControl();
  capacityControl = new FormControl();

  filteredTags: string[] = [];
  allTags: string[] = [];
  allTagsDetails: string[] = []; 

  constructor(public dialog: MatDialog, private itemService: ItemService, private adminService: AdminService) {
  }

  openDialog() {
    const dialogRef = this.dialog.open(AddRoom, {
      data: {},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        // fake data
        // this.roomsList.push({
        //   id: this.fakeId++,
        //   roomNumber: result.roomNumber,
        //   tags: result.tags,
        //   capacity: result.capacity,
        // });
        // localStorage.setItem("roomsList", JSON.stringify(this.roomsList));

        this.adminService.createRoom(
          result.roomNumber,
          result.capacity,
          result.tags.map((tag: string) => {
            const tagDetail: any = this.allTagsDetails.find((tagDetail: any) => tagDetail.tag === tag);
            return tagDetail ? tagDetail.id : null;
          }),
          ''
        ).subscribe((response: any) => {
          if (response.status === 'success') {
            this.roomsList.push({
              id: 0,
              roomNumber: result.roomNumber,
              tags: result.tags,
              capacity: result.capacity,
            });
            // localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
            console.log('Room created');}
          else{
            console.log('Room creation failed');
            return;
          }
        });

        // get all room from backend
        this.itemService.getAllRooms().subscribe((response:any)=>{
          this.roomsList = response.data.rooms.map((item: any) => {
            const tags = item.rules.map((ruleId: number) => {
              const tagDetail: any = this.allTagsDetails.find((tagDetail: any) => tagDetail.id === ruleId);
              return tagDetail ? tagDetail.tag : null;
            }).filter((tag: any) => tag !== null);
      
            return {
              id: item.id,
              roomNumber: item.room_name,
              tags: tags,
              capacity: item.capacity
            };
          });
          console.log('roomsList:', this.roomsList);

        });
        
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  ngOnInit(): void {
    // const roomsJson = localStorage.getItem("roomsList");
    // if (roomsJson) this.roomsList = JSON.parse(roomsJson);

    // get all tags from backend
    this.itemService.getAllTags().subscribe((response:any)=>{
      this.allTagsDetails = response;
      this.allTags = response.map((item: any) => item.tag );
      
    });

    // get all room from backend
    this.itemService.getAllRooms().subscribe((response:any)=>{
      this.roomsList = response.data.rooms.map((item: any) => {
        const tags = item.rules.map((ruleId: number) => {
          const tagDetail: any = this.allTagsDetails.find((tagDetail: any) => tagDetail.id === ruleId);
          return tagDetail ? tagDetail.tag : null;
        }).filter((tag: any) => tag !== null);
  
        return {
          id: item.id,
          roomNumber: item.room_name,
          tags: tags,
          capacity: item.capacity
        };
      });

    });

    
  }

  remove(tag: string, rooms: rooms): void {
    const index = rooms.tags.indexOf(tag);
    if (index >= 0) { // check if the fruit is in the list
      rooms.tags.splice(index, 1);
    }
    
  }

  selected(rooms: rooms, tag: string): void {
    if (!rooms.tags.includes(tag)) {
      rooms.tags.push(tag);
    }
  }
  
  delete(rooms: rooms): void {
    // this.roomsList = this.roomsList.filter(t => t.id !== rooms.id);
    this.adminService.deleteRoom(rooms.id).subscribe(
      (res) => {
        if (res.status === 'success') {
          this.roomsList = this.roomsList.filter(t => t.id !== rooms.id);
          console.log('Room deleted');
        }
        else{
          console.log('Delete failed');
          return
        }
      }
    );
  }

  edit(rooms: rooms): void {
    this.isEditing = !this.isEditing;
    this.roomsEditing = rooms;
    this.roomNumberControl.setValue(rooms.roomNumber);
    this.capacityControl.setValue(rooms.capacity);
    
    localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
  }

  save(): void {
    if (this.roomsEditing) {
      this.roomsEditing.roomNumber = this.roomNumberControl.value;
      this.roomsEditing.capacity = parseInt(this.capacityControl.value, 10);
      localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
    }
    this.isEditing = false;
    this.roomsEditing = undefined;
    this.roomNumberControl.setValue('');
    this.capacityControl.setValue('');
  }

}

@Component({
  selector: 'add-room',
  templateUrl: 'add-room.html',
  styleUrl: './add-room.css',
  standalone: true,
  imports: [
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatChipsModule,
    FormsModule,
    CommonModule,
    MatMenuModule,
    MatIconModule
  ],
})
export class AddRoom {
  roomNumber: string = '';
  tags: string[] = [];
  capacity: number = 0;
  allTags: string[] = [];

  constructor(
    public dialogRef: MatDialogRef<AddRoom>,
    @Inject(MAT_DIALOG_DATA) public data: any, private itemService: ItemService) {}

  ngOnInit(): void {
    // get all tags from backend
    this.itemService.getAllTags().subscribe((response:any)=>{
      this.allTags = response.map((item: any) => item.tag );
      console.log('getAllTagsArray:', this.allTags);
    });
  }

  selected(tag: string): void {
    this.tags.push(tag);
  }

  remove(tag: string): void {
    const index = this.tags.indexOf(tag);

    if (index >= 0) { // check if the fruit is in the list
      this.tags.splice(index, 1);
    }
  }

  onSave(): void {
    this.dialogRef.close({
      roomNumber: this.roomNumber,
      tags: this.tags,
      capacity: this.capacity,
    });
  }

  onCancel(): void {
    this.dialogRef.close();
  }
}