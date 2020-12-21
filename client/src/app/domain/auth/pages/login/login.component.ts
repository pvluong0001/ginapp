import {Component, OnInit} from '@angular/core';
import {AuthService} from '../../services/auth.service';
import {LoginRequest} from '../../interface';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  user: LoginRequest = {
    email: '',
    password: ''
  };

  constructor(private authService: AuthService) {
  }

  ngOnInit(): void {
  }

  handleLogin(): void {
    this.authService.loginUser(this.user)
      .subscribe(
        res => console.log(res),
        err => console.log(err)
      );
  }
}
