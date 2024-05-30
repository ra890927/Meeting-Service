import { Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map, filter } from 'rxjs/operators';


const MEETING_API = 'http://140.113.215.132:18080/api/v1/meeting/';
const USER_API = 'http://140.113.215.132:18080/api/v1/user/';
const ROOM_API = 'http://140.113.215.132:18080/api/v1/rooms/';
const TAG_API = 'http://140.113.215.132:18080/api/v1/code';
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
        console.log(response);
      const filteredData = response.data.code_types.filter((item: any) => item.type_name === 'ROOM_RULE');
      const codeValues = filteredData[0].code_values;
      console.log("codeValues", codeValues);
      return codeValues.map((item: any) => ({
        id: item.id,
        tag: item.code_value,
        description: item.code_value_desc,
        codeTypeId: item.code_type_id
        }));;
      })
    )
  }
  // get all rooms
  getAllRooms(): any {
    return this.http.get(ROOM_API, httpOptions);
  }
  // get all users
  getAllUsers(): Observable<any> {
    return this.http.get(USER_API+'getAllUsers', httpOptions);
  }
  // get meeting by user id
  getMeetingByUserId(id: string): any {
    return this.http.get(MEETING_API + "GetMeetingsByParticipantId?id=" + id, httpOptions);
  }
}
