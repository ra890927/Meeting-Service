import { CommonModule } from '@angular/common';
import { rooms } from '../users';
import { Component, Inject, OnInit } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogModule, MatDialogActions, MatDialogClose, MatDialogTitle, MatDialogContent } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { MatChipsModule } from '@angular/material/chips';
import { FormControl, FormsModule, ReactiveFormsModule, FormGroup, Validators } from '@angular/forms'
import { MatMenuModule } from '@angular/material/menu';
import { ItemService } from '../../API/item.service';
import { AdminService } from '../../API/admin.service';
import { cA } from '@fullcalendar/core/internal-common';
import { MatDividerModule } from '@angular/material/divider';

import { DeleteAlarm } from '../delete-alarm/delete-alarm';
import { MatTooltipModule } from '@angular/material/tooltip';
import { Observable, map, startWith } from 'rxjs';

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
    MatDividerModule,
    ReactiveFormsModule,
    CommonModule,
    FormsModule,
    MatTooltipModule
  ],
  templateUrl: './room.component.html',
  styleUrl: './room.component.css'
})
export class RoomComponent implements OnInit{
  
  roomsList: rooms[] = [
  //   { 
  //   id: 1,
  //   roomNumber: 'lab639',
  //   tags: ['Food Allowed'],
  //   capacity: 10,
  // },
  // { id: 2,
  //   roomNumber: 'lab637',
  //   tags: ['Food Allowed', 'Projector Available', 'Free WiFi'],
  //   capacity: 30,
  // }
];
  roomsEditing: rooms = {id: 0, roomNumber: '', tags: [], capacity: 0};
  isEditing: boolean = false;
  roomNumberControl = new FormControl();
  capacityControl = new FormControl();

  filteredTags: string[] = [];
  allTags: string[] = [];
  allTagsDetails: string[] = []; 

  filteredOptions: Observable<rooms[]> | undefined;
  roomNameSearchControl = new FormControl('');

  constructor(public dialog: MatDialog, private itemService: ItemService, private adminService: AdminService) {}

  ngOnInit(): void {
    // get all tags from backend
    this.itemService.getAllTags().subscribe((response:any)=>{
      this.allTagsDetails = response;
      console.log('allTagsDetails:', this.allTagsDetails);
      this.allTags = response.map((item: any) => item.tag );
      console.log('getAllTagsArray:', this.allTags);

    // get all room from backend
    this.itemService.getAllRooms().subscribe((response:any)=>{
      if (response.data.rooms === null) { 
        this.roomsList = [];
        return;
      }
      console.log('response:', response);
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
      console.log('checkroomsList:', this.roomsList);
      this.filteredOptions = this.roomNameSearchControl.valueChanges.pipe(
      startWith(''),
      map(value => this._filter(value || '')),
    );

    });

    });
    

    
    
  }

  private _filter(value: string): rooms[] {
    const filterValue = value.toLowerCase();
    const userRoomArray: rooms[] = this.roomsList;
    if (filterValue === '') {
      return this.roomsList;
    }
    return this.roomsList.filter(userRoomArray => userRoomArray.roomNumber.toLowerCase().includes(filterValue));
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
            console.log('Room created');

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
              this.filteredOptions = this.roomNameSearchControl.valueChanges.pipe(
                startWith(''),
                map(value => this._filter(value || '')),
              );
          
              console.log('roomsList:', this.roomsList);
            });
              
          }
          else{
            console.log('Room creation failed');
            return;
          }
        });

        
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  deleteAlarm(enterAnimationDuration: string, exitAnimationDuration: string, deleteClassType: rooms, deleteClass: string): void {
    const dialogRef = this.dialog.open(DeleteAlarm, {
      width: '250px',
      enterAnimationDuration,
      exitAnimationDuration,
      data: {deleteClassType, class: deleteClass},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result.isDelete) {
        console.log("Room delete: ", result.deleteClassType);
        this.delete(result.deleteClassType);
      }
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
          console.log('Room deleted', this.roomsList);
          this.filteredOptions = this.roomNameSearchControl.valueChanges.pipe(
            startWith(''),
            map(value => this._filter(value || '')),
          );
      
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
    console.log("roomsEditing:", this.roomsEditing);

  }

  save(): void {
    console.log('this.roomsEditing:', this.roomsEditing);
    if (this.roomsEditing) {
      console.log('Room saved', { roomNumber: this.roomNumberControl.value, capacity: this.capacityControl.value, tags: this.roomsEditing.tags });
      this.adminService.updateRoom(
        this.roomsEditing.id,
        this.roomNumberControl.value,
        this.capacityControl.value,
        this.roomsEditing.tags.map((tag: string) => {
          const tagDetail: any = this.allTagsDetails.find((tagDetail: any) => tagDetail.tag === tag);
          return tagDetail ? tagDetail.id : null;
        }),
        ''
      ).subscribe((response: any) => {
        if (response.status === 'success') {
          const index = this.roomsList.findIndex(room => room.id === this.roomsEditing.id);
          if (index !== -1) {
            this.roomsList[index].roomNumber = this.roomNumberControl.value;
            this.roomsList[index].tags = this.roomsEditing.tags;
            this.roomsList[index].capacity = this.capacityControl.value;
          }
          this.isEditing = false;
          this.roomsEditing = {id: 0, roomNumber: '', tags: [], capacity: 0};
          this.roomNumberControl.setValue('');
          this.capacityControl.setValue('');
        }
        else{
          this.isEditing = false;
          this.roomsEditing = {id: 0, roomNumber: '', tags: [], capacity: 0};
          this.roomNumberControl.setValue('');
          this.capacityControl.setValue('');
          console.log('Update failed');
          return
        }
      },(error) => {
        console.log('A connection error occurred:', error);
      }
      );
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
    MatIconModule,
    ReactiveFormsModule
  ],
})
export class AddRoom {
  roomNumber: string = '';
  tags: string[] = [];
  capacity: number = 0;
  allTags: string[] = [];
  showErrorMessage: boolean = false;
  errorMessage: string = '';
  constructor(
    public dialogRef: MatDialogRef<AddRoom>,
    @Inject(MAT_DIALOG_DATA) public data: any, private itemService: ItemService) {
    }

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
    if (this.capacity > 0 && this.roomNumber) {
      console.log('Room saved', { roomNumber: this.roomNumber, capacity: this.capacity, tags: this.tags });
      this.dialogRef.close({ roomNumber: this.roomNumber, capacity: this.capacity, tags: this.tags });
    } else {
      this.showErrorMessage = true;
      this.errorMessage = 'Please ensure the capacity is greater than 0 and the room number is filled in.';
    }
  }

  onCancel(): void {
    this.dialogRef.close();
  }
}
