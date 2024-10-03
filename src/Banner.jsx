// src/Banner.js

import React from 'react';
import bannerImage from './AcademyBanner.webp'; 
const Banner = () => {
  return (
    <div className="banner">
      <img src={bannerImage} alt="Banner" className="img-fluid" />
    </div>
  );
};

export default Banner;
