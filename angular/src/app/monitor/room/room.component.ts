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
    id: '001',
    roomNumber: 'lab639',
    tags: ['Food Allowed'],
    capacity: 10,
    details: 'This is a test room.'
  },
  { id: '002',
    roomNumber: 'lab637',
    tags: ['Food Allowed', 'Projector Available', 'Free WiFi'],
    capacity: 30,
    details: 'This is a test room2.'
  }];
  roomsEditing: rooms | undefined;
  isEditing: boolean = false;
  roomNumberControl = new FormControl();
  capacityControl = new FormControl();
  detailsControl = new FormControl();

  filteredTags: string[] = [];
  allTags = allTags;

  constructor(public dialog: MatDialog) {
  }

  openDialog() {
    const dialogRef = this.dialog.open(AddRoom, {
      data: {},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.roomsList.push({
          id: uuidv4(),
          roomNumber: result.roomNumber,
          tags: result.tags,
          capacity: result.capacity,
          details: result.details
        });
        localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
        
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  ngOnInit(): void {
    const roomsJson = localStorage.getItem("roomsList");
    if (roomsJson) this.roomsList = JSON.parse(roomsJson);
  }

  remove(tag: string, rooms: rooms): void {
    const index = rooms.tags.indexOf(tag);
    if (index >= 0) { // check if the fruit is in the list
      rooms.tags.splice(index, 1);
    }
    console.log(rooms.tags);
  }

  selected(rooms: rooms, tag: string): void {
    if (!rooms.tags.includes(tag)) {
      rooms.tags.push(tag);
    }
  }
  
  delete(rooms: rooms): void {
    this.roomsList = this.roomsList.filter(t => t.id !== rooms.id);
    localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
  }

  edit(rooms: rooms): void {
    this.isEditing = !this.isEditing;
    this.roomsEditing = rooms;
    this.roomNumberControl.setValue(rooms.roomNumber);
    this.capacityControl.setValue(rooms.capacity);
    this.detailsControl.setValue(rooms.details);
    
    localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
  }

  save(): void {
    if (this.roomsEditing) {
      this.roomsEditing.roomNumber = this.roomNumberControl.value;
      this.roomsEditing.capacity = parseInt(this.capacityControl.value, 10);
      this.roomsEditing.details = this.detailsControl.value;
      localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
    }
    this.isEditing = false;
    this.roomsEditing = undefined;
    this.roomNumberControl.setValue('');
    this.capacityControl.setValue('');
    this.detailsControl.setValue('');
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
  details: string = '';
  allTags = allTags;

  constructor(
    public dialogRef: MatDialogRef<AddRoom>,
    @Inject(MAT_DIALOG_DATA) public data: any) {}

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
        details: this.details
      });
    }

    onCancel(): void {
      this.dialogRef.close();
    }
}