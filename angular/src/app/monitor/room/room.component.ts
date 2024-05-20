import { CommonModule } from '@angular/common';
import { rooms } from '../users';
import { Component, ElementRef, ViewChild } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatInput } from '@angular/material/input';

@Component({
  selector: 'app-room',
  standalone: true,
  imports: [
    MatCardModule,
    MatListModule,
    MatIconModule,
    CommonModule
  ],
  templateUrl: './room.component.html',
  styleUrl: './room.component.css'
})
export class RoomComponent {

  @ViewChild("roomNumberInput")
  roomNumberInput!: ElementRef<MatInput>;
  @ViewChild("detailsInput")
  detailsInput!: ElementRef<MatInput>;
  
  roomsList: rooms[] = [
    { id: '001',
      roomNumber: 'lab639',
      details: 'This is a test room.'
    },
    { id: '002',
      roomNumber: 'lab637',
      details: 'This is a test room2.'
    }
  ];
  roomsEditing: rooms | undefined;
  isEditing: boolean = false;

  ngOnInit(): void {
    const roomsJson = localStorage.getItem("roomslist");
    if (roomsJson) this.roomsList = JSON.parse(roomsJson);
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
  }

}
