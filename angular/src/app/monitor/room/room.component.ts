import { CommonModule } from '@angular/common';
import { rooms } from '../users';
import { Component, ElementRef, ViewChild, Inject, OnInit, inject} from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { MatChipInputEvent, MatChipsModule } from '@angular/material/chips';
import { v4 as uuidv4 } from 'uuid'; // generate random id
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms'
import {LiveAnnouncer} from '@angular/cdk/a11y';
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
    ReactiveFormsModule,
    CommonModule,
    FormsModule,
    MatMenuModule
  ],
  templateUrl: './room.component.html',
  styleUrl: './room.component.css'
})
export class RoomComponent implements OnInit{
  @ViewChild("roomNumberInput")
  roomNumberInput!: ElementRef<MatInput>;
  @ViewChild("capacityInput")
  capacityInput!: ElementRef<MatInput>;
  @ViewChild("detailsInput")
  detailsInput!: ElementRef<MatInput>;
  
  roomsList: rooms[] = [{ 
    id: '001',
    roomNumber: 'lab639',
    fruits: ['Food Allowed'],
    capacity: 10,
    details: 'This is a test room.'
  },
  { id: '002',
    roomNumber: 'lab637',
    fruits: ['Food Allowed', 'Projector Available', 'Free WiFi'],
    capacity: 30,
    details: 'This is a test room2.'
  }];
  roomsEditing: rooms | undefined;
  isEditing: boolean = false;
  roomNumberControl = new FormControl();
  capacityControl = new FormControl();
  detailsControl = new FormControl();

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
          fruits: result.fruits,
          capacity: result.capacity,
          details: result.details
        });
        console.log(this.roomsList);
        localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
        console.log(localStorage.setItem("roomsList", JSON.stringify(this.roomsList)))
        
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  ngOnInit(): void {
    const roomsJson = localStorage.getItem("roomsList");
    if (roomsJson) this.roomsList = JSON.parse(roomsJson);
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

    setTimeout(() => {
      if (this.roomNumberInput && this.detailsInput && this.capacityInput) {
        this.roomNumberInput.nativeElement.value = rooms.roomNumber;
        this.roomNumberInput.nativeElement.focus();
        this.capacityInput.nativeElement.value = rooms.capacity;
        this.detailsInput.nativeElement.value = rooms.details;
      }
    }, 0);
    
    localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
  }

  save(): void {
    if (this.roomsEditing) {
      this.roomsEditing.roomNumber = this.roomNumberInput.nativeElement.value;
      this.roomsEditing.capacity = parseInt(this.capacityInput.nativeElement.value);
      this.roomsEditing.details = this.detailsInput.nativeElement.value;
      localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
    }
    this.isEditing = false;
    this.roomsEditing = undefined;
    this.roomNumberInput.nativeElement.value = "";
    this.capacityInput.nativeElement.value = "";
    this.detailsInput.nativeElement.value = "";
  }

  filteredTags: string[] = [];
  allTags: string[] = ['Projector Available', 'Free WiFi', 'Air Conditioning', 'Food Allowed', 'Whiteboard'];

  @ViewChild('fruitInput') fruitInput: ElementRef<HTMLInputElement> | undefined;

  announcer = inject(LiveAnnouncer);

  remove(fruit: string, rooms: rooms): void {
    const index = rooms.fruits.indexOf(fruit);

    if (index >= 0) { // check if the fruit is in the list
      rooms.fruits.splice(index, 1);

      this.announcer.announce(`Removed ${fruit}`);
    }
    console.log(rooms.fruits);
  }

  selected(event: Event, rooms: rooms, fruit: string): void {
    if (!rooms.fruits.includes(fruit)) {
      rooms.fruits.push(fruit);
    }
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
  fruits: string[] = [];
  capacity: number = 0;
  details: string = '';
  allTags: string[] = ['Projector Available', 'Free WiFi', 'Air Conditioning', 'Food Allowed', 'Whiteboard'];

  constructor(
    public dialogRef: MatDialogRef<AddRoom>,
    @Inject(MAT_DIALOG_DATA) public data: any) {}

    selected(event: Event, tag: string): void {
      this.fruits.push(tag);
    }

    remove(tag: string): void {
      const index = this.fruits.indexOf(tag);
  
      if (index >= 0) { // check if the fruit is in the list
        this.fruits.splice(index, 1);
      }
    }

    onSave(): void {
      this.dialogRef.close({
        roomNumber: this.roomNumber,
        fruits: this.fruits,
        capacity: this.capacity,
        details: this.details
      });
    }

    onCancel(): void {
      this.dialogRef.close();
    }
}