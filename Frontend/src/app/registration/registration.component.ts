import { Component, OnInit } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {
  
  username:any ='';
  password:any='';
  firstname:any='';
  lastname:any='';


  constructor(private http : HttpClient, private router: Router) { }

  ngOnInit() {
  }

  signup(){

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .post("http://34.214.86.104/signup",
        {"username":this.username,"password":this.password,"firstname":this.firstname,"lastname":this.lastname},
            {headers: header})
        .subscribe((res) => {
            //do something with the response here

            this.router.navigate(['./home']);


            console.log(res);

});
}
}
