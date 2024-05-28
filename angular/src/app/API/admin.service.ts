import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import { users } from '../monitor/users';

const USER_API = 'http://140.113.215.132:8080/api/v1/user/';
const ROOM_API = 'http://140.113.215.132:8080/api/v1/rooms/';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})

export class AdminService {

  constructor(private http: HttpClient) { 
  // admin
  // {
  //   "data": {
  //     "code_type": [
  //       {
  //         "id": 1,
  //         "type_name": "string",
  //         "type_desc": "string",
  //         "code_values": [
  //           {
  //             "id": 1,
  //             "code_type_id": 1,
  //             "code_value": "string",
  //             "code_value_desc": "string",

  //           }
  //         ]
  //       }
  //     ]
  //   },
  //  "status": "string"
  // }

  //     "message": "string",
  //     "token": "string",
  //     "user": {
  //       "created_at": "string",
  //       "email": "string",
  //       "id": 0,
  //       "role": "string",
  //       "updated_at": "string",
  //       "username": "string"
  //     }
  //   },
  //   
  // }
  }


  // getAdmins() {
  //   return this.http.get('http://localhost:3000/admins');
  // }

  updateUser(id: number, username: string, email: string, role: string): Observable<any> {
    return this.http.put(USER_API,
    {
      id,
      username,
      email,
      role
    }, 
    httpOptions);
  }

  updateRoom(id: number, room_name: string, capacity: number, rules: number[], type: string): Observable<any> {
    return this.http.put(ROOM_API,
    {
      id,
      room_name,
      capacity,
      rules,
      type
    }, 
    httpOptions);
  }

  // updateTag(id: number, tag_name: string): Observable<any> {
  // }

}
