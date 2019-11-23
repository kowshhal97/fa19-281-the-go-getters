import { Component, OnInit, Output, EventEmitter, Input } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpHeaders, HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username:any;
  password:any;
  userId:any

  requestObject ={
    "username" : this.username,
    "password": this.password
  }
  constructor(private http : HttpClient, private router: Router) {
    
  }

  
  ngOnInit() {
    

  }
  login() {

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
     this.http
        .post("https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/login",
        JSON.stringify({"username":this.username,"password":this.password}),
            {headers: header})
        .subscribe((res) => {
          sessionStorage.setItem('userId',res['id']);
          this.username=sessionStorage.setItem('username',this.username)
            //do something with the response here
            this.router.navigate(['./home']);


            console.log(res);
        }); 
    
    
  }
}