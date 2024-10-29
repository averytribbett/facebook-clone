import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'developer-page',
  templateUrl: './developer-page.component.html',
  styleUrls: ['./developer-page.component.css', '../../styles.css'],
})
export class DeveloperPageComponent {
  public isDeveloper = false;

  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.route.params.subscribe((params) => {
      this.isDeveloper = params['isDevelopmentEnvironment'];
    });
  }
}
