import { CommonModule } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';

import { tags } from '../users';
import { ItemService } from '../../API/item.service';
import { AdminService } from '../../API/admin.service';
import { co } from '@fullcalendar/core/internal-common';


@Component({
  selector: 'app-tag',
  standalone: true,
  imports: [
    MatButtonModule,
    MatCardModule,
    MatListModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule,
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    MatDialogModule,
    MatDividerModule
  ],
  templateUrl: './tag.component.html',
  styleUrl: './tag.component.css'
})
export class TagComponent {
  allTags : tags[] = [];
  codeTypeId: number = 0;
  isEditing: boolean = false;
  tagEditing = { id: 0, tag: '', description: ''};
  tagNameControl = new FormControl();
  tagDescriptionControl = new FormControl();

  fakeId = 4;
  connectionError = false;

  constructor(public dialog: MatDialog, private itemService: ItemService, private adminService: AdminService) {
  }

  ngOnInit(): void {
    // fake data
    // this.allTags.push({ id: 0, tag: 'No Smoking', description: '禁止吸菸'});
    // this.allTags.push({ id: 1, tag: 'Food Allowed', description: ''});
    // this.allTags.push({ id: 2, tag: 'Projector Available', description: ''});
    // this.allTags.push({ id: 3, tag: 'Air Conditioning', description: ''});
    // this.allTags.push({ id: 4, tag: 'Free WiFi', description: ''});
    // this.allTags.push({ id: 5, tag: 'Whiteboard', description: ''});

    // get all tags from backend
    this.itemService.getAllTags().subscribe((response:any)=>{
      this.allTags = response.map((item: any) => ({
        id: item.id,
        tag: item.tag,
        description: item.description
      }));

      this.codeTypeId = response[0].codeTypeId;
      console.log('codeTypeId:', this.codeTypeId);
      
    });
  }

  openDialog() {
    const dialogRef = this.dialog.open(AddTag, {
      data: {},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        // fake logic
        // this.allTags.push({
        //   id: this.fakeId++,
        //   tag: result.tagName,
        //   description: result.tagDescription
        // });

        // add tag to backend
        this.adminService.createTag(this.codeTypeId, result.tagName, result.tagDescription).subscribe(
          (res) => {
            if (res.status === 'success') {
              this.allTags.push({
                id: this.fakeId++,
                tag: result.tagName,
                description: result.tagDescription
              });
              console.log('Tag created');
            }
            else{
              console.log('Create failed');
              return
            }
          }
        );
        
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  delete(tag: tags): void {
    // fake logic
    // this.allTags = this.allTags.filter(t => t.id !== tag.id);
    // delete tag from backend
    this.adminService.deleteTag(tag.id).subscribe(
      (res) => {
        if (res.status === 'success') {
          this.allTags = this.allTags.filter(t => t.id !== tag.id);
          console.log('Tag deleted');
        }
        else{
          console.log('Delete failed');
          return
        }
      }
    );
  }

  edit(tag: tags): void {
    this.isEditing = !this.isEditing;
    this.tagEditing = tag;
    this.tagNameControl.setValue(tag.tag);
    this.tagDescriptionControl.setValue(tag.description);
  }

  save(): void {

    // fake logic
    //   const index = this.allTags.findIndex(tag => tag.id === this.tagEditing.id);
    //   if (index !== -1) {
    //     this.allTags[index].tag = this.tagNameControl.value;
    //     this.allTags[index].description = this.tagDescriptionControl.value;

    //   }
    // }
    // this.isEditing = false;
    // this.tagEditing = {
    //   id: 0,
    //   tag: '',
    //   description: ''
    // };
    // this.tagNameControl.setValue('');
    // this.tagDescriptionControl.setValue('');
    // console.log('allTags:', this.allTags);

    // update tag to backend
    if (this.tagEditing) {
      this.adminService.updateTag(this.codeTypeId, this.tagEditing.id, this.tagNameControl.value, this.tagDescriptionControl.value).subscribe(
        (res) => {
          if (res.status === 'success') {
            const index = this.allTags.findIndex(tag => tag.id === this.tagEditing.id);
            if (index !== -1) {
              this.allTags[index].tag = this.tagNameControl.value;
              this.allTags[index].description = this.tagDescriptionControl.value;
            }
            this.isEditing = false;
            this.tagEditing = { id: 0, tag: '', description: ''};
            this.tagNameControl.setValue('');
            this.tagDescriptionControl.setValue('');
            console.log('allTags:', this.allTags);
  
          }
          else{
            this.isEditing = false;
            this.tagEditing = { id: 0, tag: '', description: ''};
            this.tagNameControl.setValue('');
            this.tagDescriptionControl.setValue('');
            console.log('Update failed');
            return
          }
          
        },
        (error) => {
            console.error('A connection error occurred:', error);
            this.connectionError = true; 
        }
      );
    

    }

  }
}

@Component({
  selector: 'add-tag',
  templateUrl: 'add-tag.html',
  styleUrl: './add-tag.css',
  standalone: true,
  imports: [
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    FormsModule,
    CommonModule
  ],
})
export class AddTag {
  tagName: string = '';
  tagDescription: string = '';

  constructor(
    public dialogRef: MatDialogRef<AddTag>,
    @Inject(MAT_DIALOG_DATA) public data: any) {}

    onSave(): void {
      this.dialogRef.close({
        tagName: this.tagName,
        tagDescription: this.tagDescription
      });
    }

    onCancel(): void {
      this.dialogRef.close();
    }
}