import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  title = 'Login';
  toggle = true;
  books = [
    {id: 1, name: 'Book 1'},
    {id: 2, name: 'Book 2'},
  ];

  constructor() {
  }

  ngOnInit(): void {
  }

  loginHandle(): void {
    this.title = 'Logged';
  }

  getDataFromChildren(value: string): void {
    this.title = value;
  }
}
