import { Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';


const MEETING_API = 'http://140.113.215.132:8080/api/v1/meeting';
const USER_API = 'http://140.113.215.132:8080/api/v1/user/';
const ROOM_API = 'http://140.113.215.132:8080/api/v1/room/';
const TAG_API = 'http://140.113.215.132:8080/api/v1/code/type/';
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
  getAllTags(): Observable<any>  {
    return this.http.get(TAG_API + 'getAllCodeTypes', httpOptions);
  }
  // get all rooms
  getAllRooms(): Observable<any> {
    return this.http.get(ROOM_API + 'getAllRooms', httpOptions);
  }
  // get all users
  getAllUsers(): Observable<any>  {
    return this.http.get(USER_API + "getAllUsers", httpOptions);
  }
  // get meeting by user id
  getMeetingByUserId(id: string): Observable<any>  {
    return this.http.get(MEETING_API + "/GetMeetingsByParticipantId?id=" + id, httpOptions);
  }
  
  getMeetingByRoomIdAndTime(id: number, start: string, end: string): Observable<any>  {
    return this.http.get(MEETING_API + "/getMeetingsByRoomIdAndDatePeriod?room_id=" + id + "&date_from=" + start + "&date_to=" + end, httpOptions);
  }
  // post meeting
  postMeeting(description:string, end_time:string, organizer:number, participant:number[], room_id:number, start_time:string, status_type:string, title:string): Observable<any> {
    return this.http.post(MEETING_API, {
      description,
      end_time,
      organizer,
      participant,
      room_id,
      start_time,
      status_type,
      title
    }, httpOptions);
  }
  // put meeting
  putMeeting(id:string, description:string, end_time:string, organizer:number, participant:number[], room_id:number, start_time:string, status_type:string, title:string): Observable<any> {
    return this.http.put(MEETING_API, {
      id,
      description,
      end_time,
      organizer,
      participant,
      room_id,
      start_time,
      status_type,
      title
    }, httpOptions);
  }
  // delete meeting
  deleteMeeting(id:string): Observable<any> {
    return this.http.delete(MEETING_API + "/" + id, httpOptions);
  }
}
