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

  reviews:any=[]
  ngOnInit() {
    this.username=sessionStorage.getItem('username')
  }
 endpoint:any="http://13.57.219.176:80"
 itemName:any
 reviewForAnItem:any={}
 ItemNameEntered:any
 CommentsEntered:any
 RatingEntered:any
 RequestObject={}
 ReviewPost=[]
 username:any
  getReview(){

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .get(this.endpoint+'/getReviews'+'/'+this.itemName,{headers: header})
        .subscribe((res) => {
          this.reviews=res
          console.log(this.reviews)
            //do something with the response here
            this.router.navigate(['./reviews']);
            console.log(res);

});
}
postReview(){

  this.createObject()
  let header = new HttpHeaders();
  header.append('Content-Type', 'application/json');
  
   this.http
      .post(this.endpoint+'/postReview',this.RequestObject,{headers: header})
      .subscribe((res) => {
          //do something with the response here
          this.router.navigate(['./reviews']);
          console.log(res);

});
}
createObject(){
this.ReviewPost[0]={'ReviewerName':this.username,"Comment":this.CommentsEntered,"Rating":parseFloat(this.RatingEntered)}
console.log(this.ReviewPost)
this.RequestObject={"ItemName":this.ItemNameEntered,"Reviews":this.ReviewPost}
console.log(this.RequestObject)
}
}
