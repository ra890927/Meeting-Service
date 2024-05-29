import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { Component, ElementRef, ViewChild, inject,Inject} from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatAutocompleteSelectedEvent, MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatChipInputEvent, MatChipsModule} from '@angular/material/chips';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { DatePipe, AsyncPipe } from '@angular/common';
import OnInit from '@angular/core';
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
import {LiveAnnouncer} from '@angular/cdk/a11y';
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
export class PopUpFormComponent{
  separatorKeysCodes: number[] = [ENTER, COMMA];
  participantCtrl = new FormControl();
  filteredParticipants: Observable<User[]>;
  tempParticipants: User[] = [];
  tempOrganizer: User = {id: 0, username: '', email: '', role: ''};
  availableUsers:User[] = [
    {id: 1, username: 'user1', email: '1@1.com', role: 'user'},
    {id: 2, username: 'user2', email: '2@2.com', role: 'user'},
    {id: 3, username: 'user3', email: '3@3.com', role: 'user'},
    {id: 4, username: 'user4', email: '4@4.com', role: 'user'},
    {id: 5, username: 'user5', email: '5@5.com', role: 'user'},
  ];
  @ViewChild('participantInput') participantInput: ElementRef<HTMLInputElement>|undefined;
  announcer = inject(LiveAnnouncer);
  constructor(
    public dialogRef: MatDialogRef<PopUpFormComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {
    if(this.data.participants){
      this.tempParticipants = this.availableUsers.filter(user => this.data.participants.includes(user.id));
    }
    if (this.data.organizer) {
      const organizer = this.availableUsers.find(user => user.id === this.data.organizer);
  
      if (organizer) {
        this.tempOrganizer = organizer;
      } else {
        // 處理找不到組織者的情況，例如設置一個默認值或顯示錯誤訊息
        console.warn('Organizer not found');
        this.tempOrganizer = {id: 0, username: '', email: '', role: ''}; // 或者設置為其他適當的值
      }
    }
    this.filteredParticipants = this.participantCtrl.valueChanges.pipe(
      startWith(null),
      map((participant: string | null) => participant ? this._filter(participant) : this.availableUsers.slice()));
    
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
