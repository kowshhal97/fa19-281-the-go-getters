import { Component, OnInit } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {
  requestObject ={
  }

  constructor(private http : HttpClient, private router: Router) {
    
  }

  ngOnInit() {
    this.getMenu
  }

  getMenu(){
    let header=new HttpHeaders()
    header.append('content-type','application-json'),this.requestObject,
    this.http.get("endpoint/menu/items",
    {headers: header})
        .subscribe((res) => {
            //do something with the response here

            this.router.navigate(['./home']);


            console.log(res);
        });
  }
}
