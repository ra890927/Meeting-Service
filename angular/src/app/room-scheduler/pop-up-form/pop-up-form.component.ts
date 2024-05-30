import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { Component, ElementRef, ViewChild, inject,Inject} from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatAutocompleteSelectedEvent, MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatChipInputEvent, MatChipsModule} from '@angular/material/chips';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { DatePipe, AsyncPipe } from '@angular/common';
import {OnInit} from '@angular/core';
import {
  MatDialog,
  MAT_DIALOG_DATA,
  MatDialogRef,
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
  MatDialogClose,
} from '@angular/material/dialog';
import {MatButtonModule} from '@angular/material/button';
import {FormControl, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { co, ex } from '@fullcalendar/core/internal-common';
import { AuthService } from '../../API/auth.service';
import {LiveAnnouncer} from '@angular/cdk/a11y';
import { ItemService } from '../../API/item.service';
export interface DialogData {
  title: string;
  organizer: Number;//current user
  description: string;
  startTime: string;
  endTime: string;
  participants: Number[];
}
export interface User {
  id:number;
  username: string;
  email: string;
  role: string;
}
@Component({
  selector: 'app-pop-up-form',
  standalone: true,
  imports: [MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,
    DatePipe, 
    MatIconModule,
    MatSelectModule,
    CommonModule,
    MatChipsModule,
    MatAutocompleteModule,
    AsyncPipe,ReactiveFormsModule],
  templateUrl: './pop-up-form.component.html',
  styleUrl: './pop-up-form.component.css'
})
export class PopUpFormComponent implements OnInit{
  separatorKeysCodes: number[] = [ENTER, COMMA];
  participantCtrl = new FormControl();
  filteredParticipants: Observable<User[]> | undefined;
  tempParticipants: User[] = [];
  tempOrganizer: User = {id: 0, username: '', email: '', role: ''};
  availableUsers:User[] = [];
  canModify = false;
  @ViewChild('participantInput') participantInput: ElementRef<HTMLInputElement>|undefined;
  announcer = inject(LiveAnnouncer);
  constructor(
    private  auth:AuthService,
    private item: ItemService,
    public dialogRef: MatDialogRef<PopUpFormComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {
  }
  ngOnInit(): void {
    console.log("In the pop-up form component");
    this.item.getAllUsers().subscribe((response: any) => {
      this.availableUsers = response.data.users
      console.log(this.availableUsers);
      if (this.data.organizer) {
        const organizer = this.availableUsers.find(user => user.id === this.data.organizer);
    
        if (organizer) {
          this.tempOrganizer = organizer;
        } else {
          console.warn('Organizer not found');
          this.tempOrganizer = {id: 0, username: '', email: '', role: ''};
        }
        if(this.data.participants){
          this.tempParticipants = this.availableUsers.filter(user => this.data.participants.includes(user.id));
        }
        this.filteredParticipants = this.participantCtrl.valueChanges.pipe(
          startWith(null),
          map((participant: string | null) => participant ? this._filter(participant) : this.availableUsers.slice()));
      }
    });
    this.auth.whoami().subscribe((response: any) => {
      if(response.status === 'success'){
        this.canModify = this.data.organizer === response.data.user.id;
      }
    });
  }
  remove(participant: User): void {
    if(participant == this.tempOrganizer){
      this.announcer.announce('Cannot remove organizer');
      return;
    }
    const index = this.tempParticipants.indexOf(participant);
    if (index >= 0) {
      this.tempParticipants.splice(index, 1);
      this.announcer.announce('Participant removed');
    }
  }

  selected(event: MatAutocompleteSelectedEvent): void {
    this.tempParticipants.push(event.option.value);
    if(this.participantInput)
      this.participantInput.nativeElement.value = '';
    this.participantCtrl.setValue(null);
  }
  onNoClick(): void {
    this.dialogRef.close();
  }
  onOkClick(): void {
    this.data.participants = this.tempParticipants.map(participant => participant.id);
    console.log(this.data.participants);
    this.data.organizer = this.tempOrganizer.id;
  }
  private _filter(value: string): User[] {
    console.log(this.tempParticipants);
    if (typeof value !== 'string') {
      return [];
    }
    const filterValue = value.toLowerCase();
    return this.availableUsers.filter(participant => participant.username.toLowerCase().includes(filterValue) || participant.email.toLowerCase().includes(filterValue));
  }

}
