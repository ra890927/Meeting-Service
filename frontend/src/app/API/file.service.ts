import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
 
@Injectable({
  providedIn: 'root'
})
export class FileService {
  http = inject(HttpClient);
  constructor() { }

  // get file by meeting id
  getFileByMeetingId(id: string): any {
    //mock
    return this.http.get(environment.apiUrl + 'file/getFileURLsByMeetingID/' + id);
  }
  uploadFile(file: File, meeting_id: string): any {
    const formData: FormData = new FormData();
    formData.append('file', file);
    formData.append('meeting_id', String(meeting_id));
    return this.http.post(environment.apiUrl + 'file', formData);
  }
  deleteFile(file_id:string): any {
    return this.http.delete(environment.apiUrl + 'file/' + file_id);
  }
}
