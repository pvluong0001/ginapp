import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  user = {
    email: '',
    password: ''
  };

  constructor() {
  }

  ngOnInit(): void {
  }

  handleLogin(): void {
    console.log(this.user, '+++++');
  }
}