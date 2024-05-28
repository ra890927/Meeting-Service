import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { Component, ElementRef, ViewChild, inject,Inject} from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatAutocompleteSelectedEvent, MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatChipInputEvent, MatChipsModule} from '@angular/material/chips';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { DatePipe, AsyncPipe } from '@angular/common';
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
  description: string;
  startTime: string;
  endTime: string;
  participants: User[];
}
export interface User {
  name: string;
  email: string;
}
@Component({
  selector: 'app-pop-up-details',
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
  templateUrl: './pop-up-details.component.html',
  styleUrl: './pop-up-details.component.css'
})
export class PopUpDetailsComponent {
  separatorKeysCodes: number[] = [ENTER, COMMA];
  participantCtrl = new FormControl();
  filteredParticipants: Observable<User[]>;
  availableUsers:User[] = [
    {name: 'John Doe', email: 'dasdsada@gmail.com' },
    {name: 'Jane Smith', email: 'czxcxzczc@gmail.com' },
    {name: 'Bob Johnson', email: 'qweqweqweq@gmail.com' },
  ];
  tempParticipants: User[];
  @ViewChild('participantInput') participantInput: ElementRef<HTMLInputElement>|undefined;
  announcer = inject(LiveAnnouncer);
  constructor(
    public dialogRef: MatDialogRef<PopUpDetailsComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {
    this.tempParticipants = JSON.parse(JSON.stringify(data.participants));
    this.filteredParticipants = this.participantCtrl.valueChanges.pipe(
      startWith(null),
      map((participant: string | null) => participant ? this._filter(participant) : this.availableUsers.slice()));
  }
  remove(participant: User): void {
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
    console.log("in there");
    this.data.participants = this.tempParticipants;
    console.log(this.data.participants);
  }
  private _filter(value: string): User[] {
    if (typeof value !== 'string') {
      return [];
    }
    const filterValue = value.toLowerCase();
    return this.availableUsers.filter(participant => participant.name.toLowerCase().includes(filterValue) || participant.email.toLowerCase().includes(filterValue));
  }
}
