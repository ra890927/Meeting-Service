import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface File {
  id: string;
  file_name: string;
  url: string;
}
 
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
    return this.http.post('http://fake.com/api/v1/file/uploadFile', {
      file: file,
      meeting_id: meeting_id
    });
  }
  deleteFile(meeting_id: number, file_id:string): any {
    return this.http.delete('http://fake.com/api/v1/file/deleteFile?id=' + id);
  }
}
