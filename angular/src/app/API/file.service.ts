import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

 
@Injectable({
  providedIn: 'root'
})
export class FileService {
  http = inject(HttpClient);
  constructor() { }

  // get file by meeting id
  getFileByMeetingId(id: string): any {
    //mock
    return this.http.get('http://140.113.215.132:8080/api/v1/file/getFileURLsByMeetingID/' + id);
  }
  uploadFile(file: File, meeting_id: string): any {
    const formData: FormData = new FormData();
    formData.append('file', file);
    formData.append('meeting_id', String(meeting_id));
    return this.http.post('http://140.113.215.132:8080/api/v1/file', formData);
  }
  deleteFile(file_id:string): any {
    return this.http.delete('http://140.113.215.132:8080/api/v1/file/' + file_id);
  }
}
