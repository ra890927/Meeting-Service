import { CommonModule } from '@angular/common';
import { rooms } from '../users';
import { Component, ElementRef, ViewChild, Inject, OnInit} from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import {MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogModule} from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import {MatChipsModule} from '@angular/material/chips';
import { v4 as uuidv4 } from 'uuid'; // generate random id
import { FormsModule } from '@angular/forms'

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
    CommonModule
  ],
  templateUrl: './room.component.html',
  styleUrl: './room.component.css'
})
export class RoomComponent implements OnInit{

  constructor(public dialog: MatDialog) {}

  openDialog() {
    const dialogRef = this.dialog.open(AddRoom, {
      data: {},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        console.log('The dialog was closed with the following data:', result.roomNumber, result.details);
        this.roomsList.push({
          id: uuidv4(),
          roomNumber: result.roomNumber,
          tag: [
            { name: 'Projector Available', selected: false, color: 'primary' },
            { name: 'Free WiFi', selected: false, color: 'accent' },
            { name: 'Air Conditioning', selected: false, color: 'warn' }
          ],
          details: result.details
        });
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  @ViewChild("roomNumberInput")
  roomNumberInput!: ElementRef<MatInput>;
  @ViewChild("detailsInput")
  detailsInput!: ElementRef<MatInput>;
  
  roomsList: rooms[] = [
  ];
  roomsEditing: rooms | undefined;
  isEditing: boolean = false;

  ngOnInit(): void {
      this.roomsList.push({ id: '001',
      roomNumber: 'lab639',
      tag: [
        { name: 'Projector Available', selected: true, color: 'primary' },
        { name: 'Free WiFi', selected: true, color: 'primary' },
        { name: 'Air Conditioning', selected: false, color: 'primary' }
      ],
      details: 'This is a test room.'
    },
    { id: '002',
      roomNumber: 'lab637',
      tag: [
        { name: 'Projector Available', selected: false, color: 'primary' },
        { name: 'Free WiFi', selected: true, color: 'primary' },
        { name: 'Air Conditioning', selected: true, color: 'primary' }
      ],
      details: 'This is a test room2.'
    });
    localStorage.setItem("roomslist", JSON.stringify(this.roomsList));

    const roomsJson = localStorage.getItem("roomslist");
    if (roomsJson) this.roomsList = JSON.parse(roomsJson);
    console.log(this.roomsList);
  }
  
  delete(rooms: rooms): void {
    this.roomsList = this.roomsList.filter(t => t.id !== rooms.id);
    localStorage.setItem("roomslist", JSON.stringify(this.roomsList));
  }

  edit(rooms: rooms): void {
    this.isEditing = !this.isEditing;
    this.roomsEditing = rooms;

    setTimeout(() => {
      if (this.roomNumberInput && this.detailsInput) {
        this.roomNumberInput.nativeElement.value = rooms.roomNumber;
        this.roomNumberInput.nativeElement.focus();
        this.detailsInput.nativeElement.value = rooms.details;
      }
    }, 0);
    
    localStorage.setItem("roomslist", JSON.stringify(this.roomsList));
  }

  save(): void {
    if (this.roomsEditing) {
      this.roomsEditing.roomNumber = this.roomNumberInput.nativeElement.value;
      this.roomsEditing.details = this.detailsInput.nativeElement.value;
      localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
    }
    this.isEditing = false;
    this.roomsEditing = undefined;
    this.roomNumberInput.nativeElement.value = "";
    this.detailsInput.nativeElement.value = "";
  }

  add(): void {
    const roomNumber = this.roomNumberInput.nativeElement.value.trim();
    const details = this.detailsInput.nativeElement.value.trim();
    if (!roomNumber) return;
    this.roomsList.push({
      id: uuidv4(),
      roomNumber,
      tag: [
        { name: 'Projector Available', selected: false, color: 'primary' },
        { name: 'Free WiFi', selected: true, color: 'accent' },
        { name: 'Air Conditioning', selected: true, color: 'warn' }
      ],
      details

    });
    this.roomNumberInput.nativeElement.value = "";
    this.detailsInput.nativeElement.value = "";
    
    localStorage.setItem("roomslist", JSON.stringify(this.roomsList));
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
    FormsModule
  ],
})
export class AddRoom {
  roomNumber: string = '';
  details: string = '';

  constructor(
    public dialogRef: MatDialogRef<AddRoom>,
    @Inject(MAT_DIALOG_DATA) public data: any) {}

  onSave(): void {
    this.dialogRef.close({
      roomNumber: this.roomNumber,
      details: this.details
    });
  }

  onCancel(): void {
    this.dialogRef.close();
  }
}