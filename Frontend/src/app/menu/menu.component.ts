import { Component, OnInit, Input } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {

 
   menuItems:Array<any> =[
    {"_id": "5dd23a250b4efc676c0bc01b",
    "itemname": "Pepperoni Pizza",
    "itemtype": "Non-vegetarian",
    "itemsummary": "Non-veg pizza with pork",
    "itemdescription": "Pepperoni pizza with thin crust base",
    "itemamount": 10,},
    {"_id": "5dd23a250b4efc676c0bc01b",
    "itemname": "Pepperoni Pizza",
    "itemtype": "Non-vegetarian",
    "itemsummary": "Non-veg pizza with pork",
    "itemdescription": "Pepperoni pizza with thin crust base",
    "itemamount": 10,},
    {"_id": "5dd23a250b4efc676c0bc01b",
    "itemname": "Pepperoni Pizza",
    "itemtype": "Non-vegetarian",
    "itemsummary": "Non-veg pizza with pork",
    "itemdescription": "Pepperoni pizza with thin crust base",
    "itemamount": 10,}];
 
  constructor(private http : HttpClient, private router: Router) {
    
  }




  ngOnInit() {
    this.getMenu()
  }

  getMenu(){
    let header=new HttpHeaders()
    header.append('content-type','application-json')
    this.http.get<any>("endpoint/menu/items",{headers: header}).subscribe((res) => {
            //do something with the response here
            this.router.navigate(['./home']);
            this.menuItems=res
            sessionStorage.setItem('menu',res);
            console.log(res);
            
        });
  }
}
