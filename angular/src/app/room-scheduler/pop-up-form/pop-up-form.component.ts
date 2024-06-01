import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { Component, ElementRef, ViewChild, inject,Inject} from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatAutocompleteSelectedEvent, MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatChipInputEvent, MatChipsModule} from '@angular/material/chips';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { DatePipe, AsyncPipe } from '@angular/common';
import {MatMenuModule} from '@angular/material/menu';
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
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
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
import { initialEnd } from '@syncfusion/ej2-angular-schedule';
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
interface fileFormat{
  id: string;
  file_name: string;
  url: string;
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
    AsyncPipe,ReactiveFormsModule,
    MatMenuModule,
    MatProgressSpinnerModule],
  templateUrl: './pop-up-form.component.html',
  styleUrl: './pop-up-form.component.css'
})
export class PopUpFormComponent implements OnInit{
  separatorKeysCodes: number[] = [ENTER, COMMA];
  canUpload = false;
  participantCtrl = new FormControl();
  filteredParticipants: Observable<User[]> | undefined;
  uploadedFiles: fileFormat[] = [];
  tempParticipants: User[] = [];
  tempOrganizer: User = {id: 0, username: '', email: '', role: ''};
  isUploading = false;
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
        this.canUpload = this.data.participants.includes(response.data.user.id);
      }
    });
    //already have uploaded file
    
  }
  onFileSelected(e:any){
    if(this.canUpload === false){
      this.announcer.announce('You are not allowed to upload files');
      return;
    }
    this.isUploading = true;
    const files: FileList = e.target.files;
    for (let i = 0; i < files.length; i++) {
      //call the upload file api
      const file: File = files[i];
      const url = 'https://example.com/${file.name}';
      const id = '123'; // 替換為實際的上傳邏輯
      this.uploadedFiles.push({ id: id, file_name: file.name, url: url });
    }
  }
  deleteFile(file: fileFormat) {
    const index = this.uploadedFiles.indexOf(file);
    if (index > -1) {
      this.uploadedFiles.splice(index, 1);
    }
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
