import { Injectable } from '@angular/core';
import { Observable, filter, map } from 'rxjs';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import { environment } from '../../environments/environment';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})

export class AdminService {

  constructor(private http: HttpClient) {}

  updateUser(id: number, username: string, email: string, role: string, password: string): Observable<any> {
    return this.http.put(environment.apiUrl + 'admin/user',
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
    return this.http.put(environment.apiUrl + 'admin/room',
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
    return this.http.delete(environment.apiUrl + "admin/room/" + String(id), httpOptions);
  }

  createRoom(room_name: string, capacity: number, rules: number[], type: string): Observable<any> {
    return this.http.post(environment.apiUrl + 'admin/room',
      {
        room_name,
        capacity,
        rules,
        type
      }, 
      httpOptions);

  }

  updateTag(code_type_id: number, id: number, code_value: string, code_value_desc: string): Observable<any> {
    return this.http.put(environment.apiUrl + "code/value",
    {
      code_type_id,
      id,
      code_value,
      code_value_desc
    },
    httpOptions);
  }

  deleteTag(code_value_id: number): Observable<any> {
    return this.http.delete(environment.apiUrl + "code/value?id=" + String(code_value_id), httpOptions);
  }

  createTag(code_type_id: number, code_value: string, code_value_desc: string): Observable<any> {
    return this.http.post(environment.apiUrl + "code/value",
    {
      code_type_id,
      code_value,
      code_value_desc
    },
    httpOptions);
  }


}
