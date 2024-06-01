import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';

// interface File {
//   id: string;
//   file_name: string;
//   url: string;
// }
 
@Injectable({
  providedIn: 'root'
})
export class FileService {
  http = inject(HttpClient);
  constructor() { }

  // get file by meeting id
  getFileByMeetingId(id: string): any {
    //mock
    return this.http.get('http://fake.com/api/v1/file/getFileByMeetingId?id=' + id);
  }
  uploadFile(file: File, meeting_id: number): any {
    const formData: FormData = new FormData();
    formData.append('file', file);
    formData.append('file_name', file.name);
    formData.append('meeting_id', String(meeting_id));
    return this.http.post('http://fake.com/api/v1/file/uploadFile', {
      file: file,
      file_name: file.name,
      meeting_id: meeting_id
    },{
      headers: {
        'Content-Type': 'multipart/form-data'
      },
    },
  );
  }
  deleteFile(meeting_id: number, file_id:string): any {
    return this.http.delete('http://fake.com/api/v1/file/deleteFile?id=' + file_id);
  }
}
