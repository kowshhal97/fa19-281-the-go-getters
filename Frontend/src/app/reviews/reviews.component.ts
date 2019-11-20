import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpHeaders, HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.css']
})
export class ReviewsComponent implements OnInit {

  constructor(private http : HttpClient, private router: Router) { }

  ngOnInit() {
  }
 endpoint:any=""
 itemName:any=""
 reviewForAnItem:any={}
  getReview(){

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .get(this.endpoint+'/'+this.itemName,{headers: header})
        .subscribe((res) => {
          this.reviewForAnItem=res
            //do something with the response here
            this.router.navigate(['./home']);
            console.log(res);

});
}
}
