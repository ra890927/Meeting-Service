import { Injectable } from '@angular/core';
import { Observable, filter, map } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';

const USER_API = 'http://140.113.215.132:8080/api/v1/admin/user';
const ROOM_API = 'http://140.113.215.132:8080/api/v1/admin/room';
const TAG_API = 'http://140.113.215.132:8080/api/v1/code';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})

export class AdminService {

  constructor(private http: HttpClient) {}

  updateUser(id: number, username: string, email: string, role: string, password: string): Observable<any> {
    return this.http.put(USER_API,
    {
      id,
      username,
      email,
      role,
      password
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

  deleteRoom(id: number): Observable<any> {
    return this.http.delete(ROOM_API + "/" + String(id), httpOptions);
  }

  createRoom(room_name: string, capacity: number, rules: number[], type: string): Observable<any> {
    return this.http.post(ROOM_API,
      {
        room_name,
        capacity,
        rules,
        type
      }, 
      httpOptions);

  }

  updateTag(code_type_id: number, id: number, code_value: string, code_value_desc: string): Observable<any> {
    return this.http.put(TAG_API + "/value",
    {
      code_type_id,
      id,
      code_value,
      code_value_desc
    },
    httpOptions);
  }

  deleteTag(code_value_id: number): Observable<any> {
    return this.http.delete(TAG_API + "/value?id=" + String(code_value_id), httpOptions);
  }

  createTag(code_type_id: number, code_value: string, code_value_desc: string): Observable<any> {
    return this.http.post(TAG_API + "/value",
    {
      code_type_id,
      code_value,
      code_value_desc
    },
    httpOptions);
  }


}
