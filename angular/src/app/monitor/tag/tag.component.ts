import { CommonModule } from '@angular/common';
import { Component, ElementRef, Inject, ViewChild } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { allTags } from '../users';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInput, MatInputModule } from '@angular/material/input';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';


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
  allTags = allTags;
  isEditing: boolean = false;
  tagEditing: string = '';

  tagNameControl = new FormControl();

  @ViewChild("tagNameInput")
  tagNameInput!: ElementRef<MatInput>;

  constructor(public dialog: MatDialog) {
  }

  openDialog() {
    const dialogRef = this.dialog.open(AddTag, {
      data: {},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.allTags.push(result.tagName);
        
      } else {
        console.log('The dialog was closed without any data');
      }
    });
  }

  delete(tag: string): void {
    const index = allTags.indexOf(tag);
  
      if (index >= 0) { // check if the fruit is in the list
        allTags.splice(index, 1);
      }
      console.log('allTags:', allTags);
  }

  edit(tag: string): void {
    this.isEditing = !this.isEditing;
    this.tagEditing = tag;
    this.tagNameControl.setValue(tag);

  }

  save(): void {
    if (this.tagEditing) {
      const index = this.allTags.findIndex(tag => tag === this.tagEditing);

      this.tagEditing = this.tagNameInput.nativeElement.value;
      
      if (index !== -1) {
        this.allTags.splice(index, 1, this.tagNameInput.nativeElement.value);
      }
    }
    this.isEditing = false;
    this.tagEditing = '';
    this.tagNameInput.nativeElement.value = "";
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
  allTags = allTags;
  tagName: string = '';

  constructor(
    public dialogRef: MatDialogRef<AddTag>,
    @Inject(MAT_DIALOG_DATA) public data: any) {}

    onSave(): void {
      this.dialogRef.close({
        tagName: this.tagName
      });
    }

    onCancel(): void {
      this.dialogRef.close();
    }
}