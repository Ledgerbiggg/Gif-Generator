# GIF Generator

This project converts a list of images into a GIF, allowing you to resize images, set custom delays, and adjust quality settings through a simple configuration file.

## How to Use

1. **Place Your Images**: 
   - Add the images you want to use in the `images/` folder located in the root directory of the project.

2. **Configure `config.yaml`**:
   - Edit the `config.yaml` file to specify the image dimensions, file names, delays, and GIF settings.

   Example `config.yaml`:
   ```yaml
   mod: dev
   output:
     gifName: "output.gif"
   image:
     width: 800
     height: 600
     imageList:
       - file: "1.jpg"
         delay: 100
       - file: "2.jpg"
         delay: 200
       - file: "3.jpg"
         delay: 150
       - file: "4.jpg"
         delay: 100
   gifSettings:
     loop: 0
     quality: 80
   options:
     dither: true
     resizeImages: true
   ```

3. **Run the Program**:
Double click main.exe

4. **Output**:
    - The GIF will be generated as `output.gif` in the project directory.

Enjoy!

![img](https://cdn.nlark.com/yuque/0/2024/gif/35553992/1729313106459-5ced6669-3ee5-488c-8aee-cc6083791de3.gif)