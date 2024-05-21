import { CommonModule } from '@angular/common';
import { rooms } from '../users';
import { Component, ElementRef, ViewChild, Inject, OnInit} from '@angular/core';
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
    CommonModule
  ],
  templateUrl: './room.component.html',
  styleUrl: './room.component.css'
})
export class RoomComponent implements OnInit{
  @ViewChild("roomNumberInput")
  roomNumberInput!: ElementRef<MatInput>;
  @ViewChild("detailsInput")
  detailsInput!: ElementRef<MatInput>;
  
  roomsList: rooms[] = [{ 
    id: '001',
    roomNumber: 'lab639',
    tag: [
      { name: 'Projector Available', selected: true, color: 'primary' },
      { name: 'Free WiFi', selected: true, color: 'primary' },
      { name: 'Air Conditioning', selected: false, color: 'primary' } ,
      // { name: 'Projector Available1', selected: true, color: 'primary' },
      // { name: 'Free WiF2i', selected: true, color: 'primary' },
      // { name: 'Air Conditioning3', selected: false, color: 'primary' },
      // { name: 'Free WiFi4', selected: true, color: 'primary' },
      // { name: 'Air Conditioning5', selected: false, color: 'primary' } ,
      // { name: 'Projector Available6', selected: true, color: 'primary' },
      // { name: 'Free WiF7i', selected: true, color: 'primary' },
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
  }];
  roomsEditing: rooms | undefined;
  isEditing: boolean = false;
  roomNumberControl = new FormControl();
  detailsControl = new FormControl();

  constructor(public dialog: MatDialog) {}

  openDialog() {
    const dialogRef = this.dialog.open(AddRoom, {
      data: {},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.roomsList.push({
          id: uuidv4(),
          roomNumber: result.roomNumber,
          tag: result.tags,
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
    this.detailsControl.setValue(rooms.details);

    setTimeout(() => {
      if (this.roomNumberInput && this.detailsInput) {
        this.roomNumberInput.nativeElement.value = rooms.roomNumber;
        this.roomNumberInput.nativeElement.focus();
        this.detailsInput.nativeElement.value = rooms.details;
      }
    }, 0);
    
    localStorage.setItem("roomsList", JSON.stringify(this.roomsList));
  }

  toggleTagSelection(tag: { name: string, selected: boolean, color: string }) {
    tag.selected = !tag.selected;
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
    CommonModule
  ],
})
export class AddRoom {
  roomNumber: string = '';
  tag: { name: string, selected: boolean, color: string }[] = [
    { name: 'Projector Available', selected: false, color: 'primary' },
    { name: 'Free WiFi', selected: false, color: 'primary' },
    { name: 'Air Conditioning', selected: false, color: 'primary' }
  ];
  details: string = '';

  constructor(
    public dialogRef: MatDialogRef<AddRoom>,
    @Inject(MAT_DIALOG_DATA) public data: any) {}

    toggleTagSelection(tag: { name: string, selected: boolean, color: string }) {
      tag.selected = !tag.selected;
    }

    onSave(): void {
      this.dialogRef.close({
        roomNumber: this.roomNumber,
        tags: this.tag,
        details: this.details
      });
    }

    onCancel(): void {
      this.dialogRef.close();
    }
}