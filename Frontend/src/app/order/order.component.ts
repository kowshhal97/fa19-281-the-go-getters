import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { HttpHeaders, HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-order',
  templateUrl: './order.component.html',
  styleUrls: ['./order.component.css']
})
export class OrderComponent implements OnInit {


userid: any;
itemName:any
itemQuantity: any
itemPrice:any
itemId:any
getOrderId:any
deleteOrderId:any
orderId:any
  endpoint="http://34.217.1.118:3000/order"
  constructor(private http : HttpClient, private router: Router) {
    
  }
  ngOnInit() {

    this.userid =sessionStorage.getItem("userId");

  }
  placeOrder() {

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
     this.http
        .post(this.endpoint,{"userId" : this.userid,
        "itemName" : this.itemName,
        "itemQuantity" : parseInt(this.itemQuantity),
        "itemPrice": parseFloat(this.itemPrice),
        "itemId" : this.itemId},{headers: header}).subscribe((res) => {
          this.userid =sessionStorage.getItem("userId");
          console.log(this.userid)
            //do something with the response here
            this.orderId=res['orderId']
            console.log(res['orderId'])
            console.log(this.orderId)
            sessionStorage.setItem('orderId',this.orderId)
            this.router.navigate(['./payment']);
            console.log(res);
        }); 
  }
  getOrder(){
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
     this.http.get(this.endpoint+'/'+this.getOrderId,
            {headers: header})
        .subscribe((res) => {
            //do something with the response here
            this.router.navigate(['./order']);
            console.log(res);
        }); 
  }
  deleteOrder(){
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .delete(this.endpoint+'/'+this.deleteOrderId,
            {headers: header})
        .subscribe((res) => {
          this.orderId=res['orderId']
            //do something with the response here
            this.router.navigate(['./order']);
            console.log(res);
        });
  }
}

