import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpHeaders, HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-order',
  templateUrl: './order.component.html',
  styleUrls: ['./order.component.css']
})
export class OrderComponent implements OnInit {
  requestObject ={
    //user id from users
  }
  constructor(private http : HttpClient, private router: Router) {
    
  }
  ngOnInit() {

  }
  login() {

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .post("endpoint not available",
        this.requestObject,
            {headers: header})
        .subscribe((res) => {
            //do something with the response here

            this.router.navigate(['./home']);
            console.log(res);
        }); 
  }
}
