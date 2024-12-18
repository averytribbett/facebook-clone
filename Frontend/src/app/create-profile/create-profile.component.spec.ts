import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateProfileComponent } from './create-profile.component';

describe('CreateProfileComponent', () => {
  let component: CreateProfileComponent;
  let fixture: ComponentFixture<CreateProfileComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateProfileComponent],
    });
    fixture = TestBed.createComponent(CreateProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
