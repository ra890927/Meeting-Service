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
import { AsyncPipe } from '@angular/common';
import {MatAutocompleteSelectedEvent, MatAutocompleteModule} from '@angular/material/autocomplete';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import {LiveAnnouncer} from '@angular/cdk/a11y';
import { MatMenuModule } from '@angular/material/menu';
import { cellClick } from '@syncfusion/ej2-angular-schedule';

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
    MatAutocompleteModule,
    AsyncPipe,
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
    fruits: ['Food Allowed'],
    capacity: 10,
    details: 'This is a test room.'
  },
  { id: '002',
    roomNumber: 'lab637',
    tag: [
      { name: 'Projector Available', selected: false, color: 'primary' },
      { name: 'Free WiFi', selected: true, color: 'primary' },
      { name: 'Air Conditioning', selected: true, color: 'primary' }
    ],
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
    // for search filter in autocomplete
    this.filteredFruits = this.tagCtrl.valueChanges.pipe(
      startWith(null),
      map((fruit: string | null) => (fruit ? this._filter(fruit) : this.allTags.slice())),
    );

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
          tag: result.tags,
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

  toggleTagSelection(tag: { name: string, selected: boolean, color: string }) {
    tag.selected = !tag.selected;
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

  separatorKeysCodes: number[] = [ENTER, COMMA]; // for chips input search filter
  tagCtrl = new FormControl('');
  filteredFruits: Observable<string[]>;
  filteredTags: string[] = [];
  allTags: string[] = ['Projector Available', 'Free WiFi', 'Air Conditioning', 'Food Allowed', 'Whiteboard'];


  @ViewChild('fruitInput') fruitInput: ElementRef<HTMLInputElement> | undefined;

  announcer = inject(LiveAnnouncer);

  // if press separator key, then add event triggered
  add(event: MatChipInputEvent, rooms: rooms): void {
    const value = (event.value || '').trim(); // trim the unwanted spaces

    if (value && !rooms.fruits.includes(value)) {
      rooms.fruits.push(value);
      if (!this.allTags.includes(value)) {
        this.allTags.push(value);
      }
    }

    // Clear the input value
    event.chipInput!.clear();

    this.tagCtrl.setValue(null);
    console.log(rooms.fruits);
  }

  remove(fruit: string, rooms: rooms): void {
    const index = rooms.fruits.indexOf(fruit);

    if (index >= 0) { // check if the fruit is in the list
      rooms.fruits.splice(index, 1);

      this.announcer.announce(`Removed ${fruit}`);
    }
    console.log(rooms.fruits);
  }

  selected(event: MatAutocompleteSelectedEvent, rooms: rooms): void {
    
    console.log(rooms.roomNumber);
    if(!rooms.fruits.includes(event.option.viewValue)){
      rooms.fruits.push(event.option.viewValue);
    }
    if (this.fruitInput) {
      this.fruitInput.nativeElement.value = '';
    }
    this.tagCtrl.setValue(null);
    console.log(rooms.fruits);
    
  }

  selected_button(event: Event, rooms: rooms, fruit: string): void {
    if (!rooms.fruits.includes(fruit)) {
      rooms.fruits.push(fruit);
    }
  }
  

  
  // for search filter in autocomplete
  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();

    return this.allTags.filter(fruit => fruit.toLowerCase().includes(filterValue));
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
    { name: 'Air Conditioning', selected: false, color: 'primary' },
  ];
  capacity: number = 0;
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
        capacity: this.capacity,
        details: this.details
      });
    }

    onCancel(): void {
      this.dialogRef.close();
    }
}