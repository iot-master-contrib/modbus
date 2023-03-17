import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SerialComponent } from './serial.component';

describe('SerialComponent', () => {
  let component: SerialComponent;
  let fixture: ComponentFixture<SerialComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SerialComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SerialComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
