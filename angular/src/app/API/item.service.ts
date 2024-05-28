import { Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';


const MEETING_API = 'http://140.113.215.132:8080/api/v1/meeting/';
const USER_API = 'http://140.113.215.132:8080/api/v1/user/';
const ROOM_API = 'http://140.113.215.132:8080/api/v1/rooms/';
const TAG_API = 'http://140.113.215.132:8080/api/v1/tags/';
const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};
@Injectable({
  providedIn: 'root'
})
export class ItemService {

  constructor(private http: HttpClient) { 
  }
  // get all tags
  getAllTags(): any {
    return this.http.get(TAG_API, httpOptions);
  }
  // get all rooms
  getAllRooms(): any {
    return this.http.get(ROOM_API, httpOptions);
  }
  // get all users
  getAllUsers(): any {
    return this.http.get(USER_API, httpOptions);
  }
  // get meeting by user id
  getMeetingByUserId(id: string): any {
    return this.http.get(MEETING_API + "GetMeetingsByParticipantId?id=" + id, httpOptions);
  }
}
