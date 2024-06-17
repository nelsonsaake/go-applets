@REM flutter create semper_voter
@REM cmd /k cd semper_voter
cmd /k xcopy D:\desk\flutter\projects\semper_user\lib .\lib\   /S /E.
cmd /k xcopy D:\desk\flutter\projects\semper_user\assets .\assets\  /S /E.
flutter pub add stacked
flutter pub add intl
flutter pub add http
flutter pub add phosphor_flutter
flutter pub add --dev build_runner
flutter pub add --dev stacked_generator
flutter pub add stacked_services
flutter pub add responsive_builder
flutter pub add google_fonts
flutter pub add font_awesome_flutter
flutter pub add get_it
flutter pub add syncfusion_flutter_charts
flutter pub add flutter_markdown
flutter pub add flutter_svg
flutter pub add provider
flutter pub get
flutter run -d windows