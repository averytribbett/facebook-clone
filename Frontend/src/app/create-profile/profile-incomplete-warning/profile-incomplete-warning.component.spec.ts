import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileIncompleteWarningComponent } from './profile-incomplete-warning.component';

describe('ProfileIncompleteWarningComponent', () => {
  let component: ProfileIncompleteWarningComponent;
  let fixture: ComponentFixture<ProfileIncompleteWarningComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ProfileIncompleteWarningComponent]
    });
    fixture = TestBed.createComponent(ProfileIncompleteWarningComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
