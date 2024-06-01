import { Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map, filter } from 'rxjs/operators';



const MEETING_API = 'http://140.113.215.132:8080/api/v1/meeting';
const USER_API = 'http://140.113.215.132:8080/api/v1/user/';
const ROOM_API = 'http://140.113.215.132:8080/api/v1/room/';
const TAG_API = 'http://140.113.215.132:8080/api/v1/code';
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
    return this.http.get(TAG_API + "/type/getAllCodeTypes", httpOptions).pipe(
      map((response: any) => {
        // console.log(response);
      const filteredData = response.data.code_types.filter((item: any) => item.type_name === 'ROOM_RULE');
      const codeValues = filteredData[0].code_values;
      // console.log("codeValues", codeValues);
      return codeValues.map((item: any) => ({
        id: item.id,
        tag: item.code_value,
        description: item.code_value_desc,
        codeTypeId: item.code_type_id
        }));;
      })
    )

  }

  // get code type id
  getCodeTypeId(): any {
    return this.http.get(TAG_API + "/type/getAllCodeTypes", httpOptions).pipe(
      map((response: any) => {
        const filteredData = response.data.code_types.filter((item: any) => item.type_name === 'ROOM_RULE');
        return filteredData[0];
      })
    )
  }

  // get all rooms
  getAllRooms(): Observable<any> {
    return this.http.get(ROOM_API + 'getAllRooms', httpOptions);
  }
  // get all users

  getAllUsers(): Observable<any> {
    return this.http.get(USER_API+'getAllUsers', httpOptions);

  }
  // get meeting by user id
  getMeetingByUserId(id: string): Observable<any>  {
    return this.http.get(MEETING_API + "/getMeetingsByParticipantId?id=" + String(id), httpOptions);
  }
  
  getMeetingByRoomIdAndTime(id: number, start: string, end: string): Observable<any>  {
    return this.http.get(MEETING_API + "/getMeetingsByRoomIdAndDatePeriod?room_id=" + id + "&date_from=" + start + "&date_to=" + end, httpOptions);
  }
  // post meeting
  postMeeting(description:string, end_time:string, organizer:number, participants:number[], room_id:number, start_time:string, status_type:string, title:string): Observable<any> {
    return this.http.post(MEETING_API, {
      description,
      end_time,
      organizer,
      participants,
      room_id,
      start_time,
      status_type,
      title
    }, httpOptions);
  }
  // put meeting
  putMeeting(id:string, description:string, end_time:string, organizer:number, participants:number[], room_id:number, start_time:string, status_type:string, title:string): Observable<any> {
    return this.http.put(MEETING_API, {
      id,
      description,
      end_time,
      organizer,
      participants,
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
