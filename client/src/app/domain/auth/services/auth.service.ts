import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {LoginRequest} from '../interface';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private http: HttpClient) {
  }

  loginUser(payload: LoginRequest): Observable<any> {
    return this.http.post<any>('api/auth/login', payload);
  }
}
