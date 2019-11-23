import { Component, OnInit, Input } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {

 
   menuItems =[
    {"id": "1",
    "itemname": "Pepperoni Pizza",
    "itemtype": "Non-vegetarian",
    "itemsummary": "Non-veg pizza with pork",
    "itemdescription": "Pepperoni pizza with thin crust base",
    "itemamount": 10,},
    {"id": "2",
    "itemname": "Pepperoni Pizza",
    "itemtype": "Non-vegetarian",
    "itemsummary": "Non-veg pizza with pork",
    "itemdescription": "Pepperoni pizza with thin crust base",
    "itemamount": 10,},
    {"id": "5dd23a250b4efc676c0bc01b",
    "itemname": "Pepperoni Pizza",
    "itemtype": "Non-vegetarian",
    "itemsummary": "Non-veg pizza with pork",
    "itemdescription": "Pepperoni pizza with thin crust base",
    "itemamount": 10,}];
 
  constructor(private http : HttpClient, private router: Router) {
    
  }




  ngOnInit() {
    /*if(sessionStorage.getItem('userId')==null)
    {
      this.router.navigate(['./login'])
      window.alert("you need to login first!")
    }*/
    sessionStorage.setItem('menu',JSON.stringify(this.menuItems))
    this.getMenu()
    //sessionStorage.setItem('menu',JSON.stringify(this.menuItems))
  }

  getMenu(){
    let header=new HttpHeaders()
    header.append('content-type','application-json')
    this.http.get<any>("https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/menu",{headers: header}).subscribe((res) => {
            //do something with the response here
            this.router.navigate(['./menu']);
            this.menuItems=res
            sessionStorage.setItem('menu',res);
            console.log(res);
            
        });
  }
  gotoMenu(){
    this.router.navigate(['./menu'])
  }

  gotoReviews(){
    this.router.navigate(['./reviews'])
  }

  gotoHome(){
    this.router.navigate(['./home'])
  }
  logout(){
    sessionStorage.setItem('userId',null)
    this.router.navigate(['./login'])
  }
}
