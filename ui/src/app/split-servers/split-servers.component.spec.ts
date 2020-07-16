/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { SplitServersComponent } from './split-servers.component';

describe('SplitServersComponent', () => {
  let component: SplitServersComponent;
  let fixture: ComponentFixture<SplitServersComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SplitServersComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SplitServersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
