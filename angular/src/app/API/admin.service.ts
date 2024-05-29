import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';

const USER_API = 'http://140.113.215.132:8080/api/v1/admin/user';
const ROOM_API = 'http://140.113.215.132:8080/api/v1/admin/rooms';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})

export class AdminService {

  constructor(private http: HttpClient) {}


  // getAdmins() {
  //   return this.http.get('http://localhost:3000/admins');
  // }

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
    return this.http.delete(ROOM_API + id.toString(), httpOptions);
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


  // updateTag(id: number, tag_name: string): Observable<any> {
  // }

}
