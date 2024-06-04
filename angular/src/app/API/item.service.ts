import { Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map, filter } from 'rxjs/operators';
import { environment } from '../../environments/environment';

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
    return this.http.get(environment.apiUrl + "code/type/getAllCodeTypes", httpOptions).pipe(
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
    return this.http.get(environment.apiUrl + 'room/getAllRooms', httpOptions);
  }
  // get all users

  getAllUsers(): Observable<any> {
    return this.http.get(environment.apiUrl + 'user/getAllUsers', httpOptions);

  }
  // get meeting by user id
  getMeetingByUserId(id: string): Observable<any>  {
    return this.http.get(environment.apiUrl + "room/getMeetingsByParticipantId?id=" + String(id), httpOptions);
  }
  
  getMeetingByRoomIdAndTime(id: number, start: string, end: string): Observable<any>  {
    return this.http.get(environment.apiUrl + "meeting/getMeetingsByRoomIdAndDatePeriod?room_id=" + id + "&date_from=" + start + "&date_to=" + end, httpOptions);
  }
  // post meeting
  postMeeting(description:string, end_time:string, organizer:number, participants:number[], room_id:number, start_time:string, status_type:string, title:string): Observable<any> {
    return this.http.post(environment.apiUrl + 'meeting', {
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
    return this.http.put(environment.apiUrl + 'meeting', {
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
    return this.http.delete(environment.apiUrl + 'meeting/' + id, httpOptions);
  }
}
