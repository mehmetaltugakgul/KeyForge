# KeyForge: Strong Password Generator

KeyForge is a desktop application designed to create strong and customizable passwords.

<img width="338" alt="image" src="https://github.com/user-attachments/assets/06034e05-5058-47f1-89a9-7b7581ed09ae">


## Features

- Adjustable password length (4-128 characters)
- Options for uppercase letters, lowercase letters, numbers, and special characters
- Option to include at least one of each selected character type
- Password strength indicator
- Copy generated password to clipboard
- System tray support

## Installation

1. Download and install Go from the [official website](https://golang.org/).

2. Install the Fyne library:
   ```
   go get fyne.io/fyne/v2
   ```
3. Clone this project:
   ```
   git clone https://github.com/mehmetaltugakgul/keyforge.git
   ```
4. Navigate to the project directory:
   ```
   cd keyforge
   ```
5. Compile and run the application:
   ```
   go run main.go
   ```

## Usage

1. Launch the application.
2. Adjust the desired password length using the slider.
3. Select the character types you want to include in your password.
4. Check the "At Least One of Each Selected Type" option to ensure the password contains at least one character from each selected type.
5. Click the "Generate Password" button to create your password.
6. View the generated password and check its strength.
7. Use the "Copy Password" button to copy the password to your clipboard.

## Building

To build the application for your current platform:

## Contributing

If you'd like to contribute to the project, please fork the repository and create a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the [MIT License](LICENSE).
